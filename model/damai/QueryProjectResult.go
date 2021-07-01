package damai

// QueryProjectResult 
type QueryProjectResult struct {
    // 项目信息
    Projects   []ProjectDto `json:"projects,omitempty" xml:"projects>project_dto,omitempty"`
}
