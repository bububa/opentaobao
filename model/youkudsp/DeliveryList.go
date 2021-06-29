package youkudsp

// DeliveryList 
type DeliveryList struct {

    // 投放类型push或者feed
    
    DeliveryType   string `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
    

    // 渠道id
    
    ChannelId   int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
    

    // 素材信息
    
    Resource   *YoukuDspDeliveryResourceMultigetMap `json:"resource,omitempty" xml:"resource,omitempty"`
    

    // 设备类型imei或者idfa
    
    DeviceIdType   string `json:"device_id_type,omitempty" xml:"device_id_type,omitempty"`
    

    // 子渠道id
    
    SubChannelId   int64 `json:"sub_channel_id,omitempty" xml:"sub_channel_id,omitempty"`
    

    // 设备id(md5加密)
    
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    

}
