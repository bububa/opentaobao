package admarket

// BidRequest 
type BidRequest struct {

    // 应用信息
    
    AppInfo   *AppInfo `json:"app_info,omitempty" xml:"app_info,omitempty"`
    

    // sdk信息
    
    SdkInfo   *SdkInfo `json:"sdk_info,omitempty" xml:"sdk_info,omitempty"`
    

    // 设备信息
    
    DeviceInfo   *DeviceInfo `json:"device_info,omitempty" xml:"device_info,omitempty"`
    

    // 设备id
    
    Udid   *Udid `json:"udid,omitempty" xml:"udid,omitempty"`
    

    // 用户信息
    
    UserInfo   *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
    

    // 渠道id
    
    Channel   string `json:"channel,omitempty" xml:"channel,omitempty"`
    

    // 网络信息
    
    Network   *Network `json:"network,omitempty" xml:"network,omitempty"`
    

    // 版本号
    
    Version   string `json:"version,omitempty" xml:"version,omitempty"`
    

    // 定位信息
    
    Location   *Location `json:"location,omitempty" xml:"location,omitempty"`
    

    // 广告位列表
    
    AdSlots   []AdSlot `json:"ad_slots,omitempty" xml:"ad_slots,omitempty"`
    

    // 批次id
    
    BatchId   string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
    

}
