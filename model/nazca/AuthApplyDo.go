package nazca

// AuthApplyDo 结构体
type AuthApplyDo struct {
	// 企业名称
	EnterpriseName string `json:"enterprise_name,omitempty" xml:"enterprise_name,omitempty"`
	// 营业执照号
	License string `json:"license,omitempty" xml:"license,omitempty"`
	// 组织机构代码
	Organization string `json:"organization,omitempty" xml:"organization,omitempty"`
	// 客户在1688的唯一标识
	PlatformUserId string `json:"platform_user_id,omitempty" xml:"platform_user_id,omitempty"`
	// 页面跳转同步通知URL
	ReturnUrl string `json:"return_url,omitempty" xml:"return_url,omitempty"`
	// 认证状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 统一社会信用代码
	ThreeCertNumber string `json:"three_cert_number,omitempty" xml:"three_cert_number,omitempty"`
	// 证件类型 1：旧版三证 2：新版三证合一
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 是否认证
	Autherized bool `json:"autherized,omitempty" xml:"autherized,omitempty"`
}
