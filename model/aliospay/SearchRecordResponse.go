package aliospay

// SearchRecordResponse 
type SearchRecordResponse struct {

    // 总数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

    // 支付记录列表
    
    Datas   []PayRecordData `json:"datas,omitempty" xml:"datas,omitempty"`
    

}
