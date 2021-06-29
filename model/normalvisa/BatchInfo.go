package normalvisa

// BatchInfo 
type BatchInfo struct {

    // 批次id
    
    BatchId   string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
    

    // 每批次的申请人id
    
    ApplyIds   []string `json:"apply_ids,omitempty" xml:"apply_ids>string,omitempty"`
    

}
