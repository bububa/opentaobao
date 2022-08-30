package util

// AiotTopOpenDeviceDto 结构体
type AiotTopOpenDeviceDto struct {
	// 设备唯一标识
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 业务自定义参数
	ExtendStr string `json:"extend_str,omitempty" xml:"extend_str,omitempty"`
	// 组织架构
	Organization *AiotOpenDeviceOrganizationDto `json:"organization,omitempty" xml:"organization,omitempty"`
	// 基础信息
	Base *AiotOpenDeviceBaseDto `json:"base,omitempty" xml:"base,omitempty"`
	// 租户号
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}
