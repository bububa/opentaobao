package tmallsc

// IdentifyTaskDeliveryRequest 
type IdentifyTaskDeliveryRequest struct {

    // 工单ID
    
    WorkcardId   int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
    

    // 配送地纬度
    
    Latitude   string `json:"latitude,omitempty" xml:"latitude,omitempty"`
    

    // 配送地经度
    
    Longitude   string `json:"longitude,omitempty" xml:"longitude,omitempty"`
    

}
