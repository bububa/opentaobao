package mtopopen

// GeneralLogisticsDataWriteRequest 结构体
type GeneralLogisticsDataWriteRequest struct {
	// 快递公司标准编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 运单关联包裹的当前物流状态，需和平台侧定义的枚举名保持一致。
	CurrentLogisticsStatus string `json:"current_logistics_status,omitempty" xml:"current_logistics_status,omitempty"`
	// 运单关联包裹的收件人手机号。淘内包裹必须不填，淘外包裹必填，否则无法发送push消息
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 待写入的数据，key值需要按照平台侧定义的标准。isSupportDeliveryHome：是否支持送货上门。hasSignPic：是否有签收照片。value为String类型，TRUE代表真，FALSE代表假
	ExtendParam string `json:"extend_param,omitempty" xml:"extend_param,omitempty"`
	// 用户在手淘中登陆并授权小程序后的唯一标识。登陆过淘宝小程序的用户都会有一个openid，如果有尽量传过来，因为如果缺少openid，则无法判断用户是否为高活用户，也就无法决策是否推送push消息
	Openid string `json:"openid,omitempty" xml:"openid,omitempty"`
	// INTERNAL：淘内包裹，EXTERNAL：淘外包裹
	PackageType string `json:"package_type,omitempty" xml:"package_type,omitempty"`
	// DELIVERY_HOME：送货上门。SIGN_PIC：签收照片
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
}
