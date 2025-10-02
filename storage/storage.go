package storage

import (
	"dev-util/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	configDir  = ".dev-util"
	configFile = "projects.json"
)

// GetConfigPath returns the path to the configuration file
func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	
	configDirPath := filepath.Join(homeDir, configDir)
	if err := os.MkdirAll(configDirPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create config directory: %w", err)
	}
	
	return filepath.Join(configDirPath, configFile), nil
}

// LoadProjects loads projects from the configuration file
func LoadProjects() (*models.ProjectStore, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}
	
	// If file doesn't exist, return empty store
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &models.ProjectStore{Projects: []models.Project{}}, nil
	}
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	
	var store models.ProjectStore
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	
	return &store, nil
}

// SaveProjects saves projects to the configuration file
func SaveProjects(store *models.ProjectStore) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}
	
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal projects: %w", err)
	}
	
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	
	return nil
}

// AddProject adds a new project and saves it
func AddProject(name, path, command, description string) error {
	store, err := LoadProjects()
	if err != nil {
		return err
	}
	
	// Check if project already exists
	if _, exists := store.GetProject(name); exists {
		return fmt.Errorf("project '%s' already exists", name)
	}
	
	project := models.Project{
		Name:        name,
		Path:        path,
		Command:     command,
		Description: description,
		CreatedAt:   time.Now(),
	}
	
	store.AddProject(project)
	return SaveProjects(store)
}

// GetProject retrieves a project by name
func GetProject(name string) (*models.Project, error) {
	store, err := LoadProjects()
	if err != nil {
		return nil, err
	}
	
	project, exists := store.GetProject(name)
	if !exists {
		return nil, fmt.Errorf("project '%s' not found", name)
	}
	
	return project, nil
}

// ListProjects returns all projects
func ListProjects() ([]models.Project, error) {
	store, err := LoadProjects()
	if err != nil {
		return nil, err
	}
	
	return store.ListProjects(), nil
}

// RemoveProject removes a project by name
func RemoveProject(name string) error {
	store, err := LoadProjects()
	if err != nil {
		return err
	}
	
	if !store.RemoveProject(name) {
		return fmt.Errorf("project '%s' not found", name)
	}
	
	return SaveProjects(store)
}
