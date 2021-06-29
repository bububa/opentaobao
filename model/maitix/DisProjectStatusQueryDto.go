package maitix

// DisProjectStatusQueryDTO 
type DisProjectStatusQueryDTO struct {
    // 是否查询对应场次的状态
    QueryPerformStatus   bool `json:"query_perform_status,omitempty" xml:"query_perform_status,omitempty"`
    // 项目id
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
}
