package maitix

// DisPerformStatusDTO 
type DisPerformStatusDTO struct {
    // 场次id
    PerformId   int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
    // 项目id
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
    // 场次状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}
