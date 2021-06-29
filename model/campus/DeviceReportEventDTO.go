package campus

// DeviceReportEventDTO 
type DeviceReportEventDTO struct {

    // 0:设备中心UUID ，1：外部id
    
    IdType   int64 `json:"id_type,omitempty" xml:"id_type,omitempty"`
    

    // 数据
    
    Data   []null `json:"data,omitempty" xml:"data,omitempty"`
    

    // 消息唯一id
    
    TransId   string `json:"trans_id,omitempty" xml:"trans_id,omitempty"`
    

    // 消息时间戳
    
    EventTime   int64 `json:"event_time,omitempty" xml:"event_time,omitempty"`
    

    // 来源系统
    
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    

    // 消息版本
    
    Version   string `json:"version,omitempty" xml:"version,omitempty"`
    

    // 设备ID
    
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    

}
