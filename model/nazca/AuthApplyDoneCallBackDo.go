package nazca

// AuthApplyDoneCallBackDo 结构体
type AuthApplyDoneCallBackDo struct {
	// 身份证
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 支付宝账号
	AlipayAccount string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
	// 支付宝用户id
	AlipayUserId string `json:"alipay_user_id,omitempty" xml:"alipay_user_id,omitempty"`
	// 申请人身份： 1：企业法人/负责人；2：企业代理人
	ApplicantType string `json:"applicant_type,omitempty" xml:"applicant_type,omitempty"`
	// 证书DN
	Dn string `json:"dn,omitempty" xml:"dn,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 证书有效期截止时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 待认证的企业名称
	EnterpriseName string `json:"enterprise_name,omitempty" xml:"enterprise_name,omitempty"`
	// 营业执照号
	License string `json:"license,omitempty" xml:"license,omitempty"`
	// 组织机构代码
	Organization string `json:"organization,omitempty" xml:"organization,omitempty"`
	// 对应人员名称
	PersonName string `json:"person_name,omitempty" xml:"person_name,omitempty"`
	// 证书pfx文件
	Pfx string `json:"pfx,omitempty" xml:"pfx,omitempty"`
	// 证书pfx密钥
	PfxKey string `json:"pfx_key,omitempty" xml:"pfx_key,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 客户在1688的唯一标识
	PlatformUserId string `json:"platform_user_id,omitempty" xml:"platform_user_id,omitempty"`
	// 证书序列号
	SequenceNo string `json:"sequence_no,omitempty" xml:"sequence_no,omitempty"`
	// 证书有效期起始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 审核状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 统一社会信用代码
	ThreeCertNumber string `json:"three_cert_number,omitempty" xml:"three_cert_number,omitempty"`
	// 证件类型 1：旧版三证 2：新版三证合一
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
