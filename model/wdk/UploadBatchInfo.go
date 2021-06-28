package wdk

// UploadBatchInfo 
/* model for simplify = false
type UploadBatchInfo struct {

    // 批次号
    
    BatchNo   string `json:"batch_no,omitempty"`
    

    // 批次总量
    
    BatchCount   int64 `json:"batch_count,omitempty"`
    

    // 该批次数据的时间范围-起始时间
    
    BeginTime   string `json:"begin_time,omitempty"`
    

    // 该批次数据的时间范围-结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

}
*/

// UploadBatchInfo 
type UploadBatchInfo struct {

    // 批次号
    BatchNo   string `json:"batch_no,omitempty"`

    // 批次总量
    BatchCount   int64 `json:"batch_count,omitempty"`

    // 该批次数据的时间范围-起始时间
    BeginTime   string `json:"begin_time,omitempty"`

    // 该批次数据的时间范围-结束时间
    EndTime   string `json:"end_time,omitempty"`

}
