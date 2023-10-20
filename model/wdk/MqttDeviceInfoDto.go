package wdk

import (
	"sync"
)

// MqttDeviceInfoDto 结构体
type MqttDeviceInfoDto struct {
	// mqtt设备名
	DeviceName string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	// mqtt设备秘钥
	DeviceSecret string `json:"device_secret,omitempty" xml:"device_secret,omitempty"`
	// mqtt设备所属产品
	ProductKey string `json:"product_key,omitempty" xml:"product_key,omitempty"`
	// accessKey
	AccessKey string `json:"access_key,omitempty" xml:"access_key,omitempty"`
	// accessToken
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// expireTime
	ExpireTime int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// 创建令牌的服务端时间戳
	Timestamp int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

var poolMqttDeviceInfoDto = sync.Pool{
	New: func() any {
		return new(MqttDeviceInfoDto)
	},
}

// GetMqttDeviceInfoDto() 从对象池中获取MqttDeviceInfoDto
func GetMqttDeviceInfoDto() *MqttDeviceInfoDto {
	return poolMqttDeviceInfoDto.Get().(*MqttDeviceInfoDto)
}

// ReleaseMqttDeviceInfoDto 释放MqttDeviceInfoDto
func ReleaseMqttDeviceInfoDto(v *MqttDeviceInfoDto) {
	v.DeviceName = ""
	v.DeviceSecret = ""
	v.ProductKey = ""
	v.AccessKey = ""
	v.AccessToken = ""
	v.ExpireTime = 0
	v.Timestamp = 0
	poolMqttDeviceInfoDto.Put(v)
}
