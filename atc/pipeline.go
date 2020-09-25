package atc

type Pipeline struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Paused      bool           `json:"paused"`
	Public      bool           `json:"public"`
	Archived    bool           `json:"archived"`
	Groups      GroupConfigs   `json:"groups,omitempty"`
	TeamName    string         `json:"team_name"`
	Display     *DisplayConfig `json:"display,omitempty"`
	LastUpdated int64          `json:"last_updated,omitempty"`
}

type KubernetesPipeline struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`
	Spec       PipelineSpec   `json:"spec,omitempty"`
	Status     PipelineStatus `json:"status,omitempty"`
}

type PipelineSpec struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	TeamName string `json:"team_name"`
}

type PipelineStatus struct {
	Paused      bool           `json:"paused"`
	Public      bool           `json:"public"`
	Archived    bool           `json:"archived"`
	Groups      GroupConfigs   `json:"groups,omitempty"`
	Display     *DisplayConfig `json:"display,omitempty"`
	LastUpdated int64          `json:"last_updated,omitempty"`
}

type RenameRequest struct {
	NewName string `json:"name"`
}
