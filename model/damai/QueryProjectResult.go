package damai

// QueryProjectResult 
type QueryProjectResult struct {
    // 项目信息
    Projects   []ProjectDTO `json:"projects,omitempty" xml:"projects>project_dto,omitempty"`
}
