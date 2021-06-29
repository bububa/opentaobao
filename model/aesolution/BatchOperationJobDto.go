package aesolution

// BatchOperationJobDto 
type BatchOperationJobDto struct {
    // The status of the feed
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // feed type
    OperationType   string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
    // job id
    JobId   int64 `json:"job_id,omitempty" xml:"job_id,omitempty"`
}
