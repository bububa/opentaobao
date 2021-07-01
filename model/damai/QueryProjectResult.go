package damai

// QueryProjectResult 结构体
type QueryProjectResult struct {
	// 项目信息
	Projects []ProjectDto `json:"projects,omitempty" xml:"projects>project_dto,omitempty"`
}
