package cmns

// DeviceResult 
type DeviceResult struct {

    // 是否接收消息
    AcceptMessage   int64 `json:"accept_message,omitempty"`

    // 激活时间
    ActiveTime   string `json:"active_time,omitempty"`

    // 基带
    BaseType   string `json:"base_type,omitempty"`

    // BSP
    BspType   string `json:"bsp_type,omitempty"`

    // 设备类型：1手机 2TV
    DeviceType   int64 `json:"device_type,omitempty"`

    // 有线网卡mac
    EthMac   string `json:"eth_mac,omitempty"`

    // did
    Id   int64 `json:"id,omitempty"`

    // imei
    Imei   string `json:"imei,omitempty"`

    // imsi1
    Imsi1   string `json:"imsi1,omitempty"`

    // imsi2
    Imsi2   string `json:"imsi2,omitempty"`

    // isVendorRom
    IsVendorRom   string `json:"is_vendor_rom,omitempty"`

    // kp
    Kp   string `json:"kp,omitempty"`

    // 最近登录时间
    LoginTime   int64 `json:"login_time,omitempty"`

    // 网络类型
    Networking   string `json:"networking,omitempty"`

    // 是否在线,1为在线，0为离线
    Online   int64 `json:"online,omitempty"`

    // phoneType
    PhoneType   string `json:"phone_type,omitempty"`

    // 系统版本
    Rom   string `json:"rom,omitempty"`

    // romTuiguang
    RomTuiguang   string `json:"rom_tuiguang,omitempty"`

    // romVendor
    RomVendor   string `json:"rom_vendor,omitempty"`

    // 设备机型
    Terminal   string `json:"terminal,omitempty"`

    // 轻量升级版本
    UpdateVersion   string `json:"update_version,omitempty"`

    // uuid
    Uuid   string `json:"uuid,omitempty"`

    // 无线mac
    WlanMac   string `json:"wlan_mac,omitempty"`

}
