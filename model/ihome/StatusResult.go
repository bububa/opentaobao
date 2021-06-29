package ihome

// StatusResult 
type StatusResult struct {
    // 状态list
    StatusList   []ContentStatus `json:"status_list,omitempty" xml:"status_list>content_status,omitempty"`
}
