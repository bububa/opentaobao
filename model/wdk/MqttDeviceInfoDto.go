package wdk

// MqttDeviceInfoDto 
/* model for simplify = false
type MqttDeviceInfoDto struct {

    // mqtt设备名
    
    DeviceName   string `json:"device_name,omitempty"`
    

    // mqtt设备秘钥
    
    DeviceSecret   string `json:"device_secret,omitempty"`
    

    // mqtt设备所属产品
    
    ProductKey   string `json:"product_key,omitempty"`
    

    // accessKey
    
    AccessKey   string `json:"access_key,omitempty"`
    

    // accessToken
    
    AccessToken   string `json:"access_token,omitempty"`
    

    // expireTime
    
    ExpireTime   int64 `json:"expire_time,omitempty"`
    

    // 创建令牌的服务端时间戳
    
    Timestamp   int64 `json:"timestamp,omitempty"`
    

}
*/

// MqttDeviceInfoDto 
type MqttDeviceInfoDto struct {

    // mqtt设备名
    DeviceName   string `json:"device_name,omitempty"`

    // mqtt设备秘钥
    DeviceSecret   string `json:"device_secret,omitempty"`

    // mqtt设备所属产品
    ProductKey   string `json:"product_key,omitempty"`

    // accessKey
    AccessKey   string `json:"access_key,omitempty"`

    // accessToken
    AccessToken   string `json:"access_token,omitempty"`

    // expireTime
    ExpireTime   int64 `json:"expire_time,omitempty"`

    // 创建令牌的服务端时间戳
    Timestamp   int64 `json:"timestamp,omitempty"`

}
