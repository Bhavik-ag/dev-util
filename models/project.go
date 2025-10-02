package models

import "time"

// Project represents a development project configuration
type Project struct {
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Command     string    `json:"command"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// ProjectStore manages the collection of projects
type ProjectStore struct {
	Projects []Project `json:"projects"`
}

// AddProject adds a new project to the store
func (ps *ProjectStore) AddProject(project Project) {
	ps.Projects = append(ps.Projects, project)
}

// GetProject retrieves a project by name
func (ps *ProjectStore) GetProject(name string) (*Project, bool) {
	for _, project := range ps.Projects {
		if project.Name == name {
			return &project, true
		}
	}
	return nil, false
}

// ListProjects returns all projects
func (ps *ProjectStore) ListProjects() []Project {
	return ps.Projects
}

// RemoveProject removes a project by name
func (ps *ProjectStore) RemoveProject(name string) bool {
	for i, project := range ps.Projects {
		if project.Name == name {
			ps.Projects = append(ps.Projects[:i], ps.Projects[i+1:]...)
			return true
		}
	}
	return false
}
