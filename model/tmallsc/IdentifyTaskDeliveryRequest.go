package tmallsc

// IdentifyTaskDeliveryRequest 
/* model for simplify = false
type IdentifyTaskDeliveryRequest struct {

    // 工单ID
    
    WorkcardId   int64 `json:"workcard_id,omitempty"`
    

    // 配送地纬度
    
    Latitude   string `json:"latitude,omitempty"`
    

    // 配送地经度
    
    Longitude   string `json:"longitude,omitempty"`
    

}
*/

// IdentifyTaskDeliveryRequest 
type IdentifyTaskDeliveryRequest struct {

    // 工单ID
    WorkcardId   int64 `json:"workcard_id,omitempty"`

    // 配送地纬度
    Latitude   string `json:"latitude,omitempty"`

    // 配送地经度
    Longitude   string `json:"longitude,omitempty"`

}
