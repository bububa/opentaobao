package ascp

import (
	"sync"
)

// DeliveryInfo 结构体
type DeliveryInfo struct {
	// 商家编码，商家在erp维护的编码
	ErpDeliveryBizCode string `json:"erp_delivery_biz_code,omitempty" xml:"erp_delivery_biz_code,omitempty"`
	// 资源名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 平台资源编码，枚举：宅急送=ZJS、顺丰速运=SF、申通快递=STO、EMS=EMS、韵达快递=YUNDA、极兔速递=HTKY、圆通速递=YTO、天天快递=TTKDEX、全峰快递=QFKD、EMS经济快递=EYB、优速快递=UC、德邦快递=DBKD、速尔快运=SURE、联邦快递=FEDEX、华强物流=SHQ、全一快递=UAPEX、天地华宇=HOAU、百世物流=BEST、龙邦速递=LB、新邦物流=XB、中通快递=ZTO、国通快递=GTO、快捷快=FAST、能达速递=NEDA、如风达配送=BJRFD-001、信丰物流=XFWL、广东EMS=GDEMS、邮政快递包裹=POSTB、德邦物流=DBL、黑猫宅急便=YCTE速宝=ESB、联昊通= LTS、E速宝=ESB、百世快运=BESTQJT、增益速递=QRT、其他=OTHERS
	PlatformCode string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
	// 联系人姓名
	ConName string `json:"con_name,omitempty" xml:"con_name,omitempty"`
	// 联系人电话
	ConPhone string `json:"con_phone,omitempty" xml:"con_phone,omitempty"`
	// erp配资源唯一编码，卖家唯一
	ErpCode string `json:"erp_code,omitempty" xml:"erp_code,omitempty"`
	// 状态：0=停用；1=启用
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolDeliveryInfo = sync.Pool{
	New: func() any {
		return new(DeliveryInfo)
	},
}

// GetDeliveryInfo() 从对象池中获取DeliveryInfo
func GetDeliveryInfo() *DeliveryInfo {
	return poolDeliveryInfo.Get().(*DeliveryInfo)
}

// ReleaseDeliveryInfo 释放DeliveryInfo
func ReleaseDeliveryInfo(v *DeliveryInfo) {
	v.ErpDeliveryBizCode = ""
	v.Name = ""
	v.PlatformCode = ""
	v.ConName = ""
	v.ConPhone = ""
	v.ErpCode = ""
	v.Status = ""
	poolDeliveryInfo.Put(v)
}
