package logistic

import (
	"sync"
)

// Param 结构体
type Param struct {
	// 门店编码，对应大润发deliveryDockCode
	StoreCodes []string `json:"store_codes,omitempty" xml:"store_codes>string,omitempty"`
	// 门店列表
	Chainstores []ChainStore `json:"chainstores,omitempty" xml:"chainstores>chain_store,omitempty"`
	// appId
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 外部订单号
	PartnerOrderCode string `json:"partner_order_code,omitempty" xml:"partner_order_code,omitempty"`
	// 渠道code
	ChannelId string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 取消原因描述
	OrderCancelDesc string `json:"order_cancel_desc,omitempty" xml:"order_cancel_desc,omitempty"`
	// 服务包code
	ServicePackageCode string `json:"service_package_code,omitempty" xml:"service_package_code,omitempty"`
	// 商户code
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 门店code
	ChainstoreCode string `json:"chainstore_code,omitempty" xml:"chainstore_code,omitempty"`
	// 文本地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 门店名字
	ChainStoreName string `json:"chain_store_name,omitempty" xml:"chain_store_name,omitempty"`
	// 门店code
	ChainStoreCode string `json:"chain_store_code,omitempty" xml:"chain_store_code,omitempty"`
	// 联系手机
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 错误详情
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 取消code
	OrderCancelCode int64 `json:"order_cancel_code,omitempty" xml:"order_cancel_code,omitempty"`
	// 取消原因code  1:用户取消
	OrderCancelReasonCode int64 `json:"order_cancel_reason_code,omitempty" xml:"order_cancel_reason_code,omitempty"`
	// 经纬度来源。1-腾讯，2-百度，3-高德
	PositionSource int64 `json:"position_source,omitempty" xml:"position_source,omitempty"`
	// 交易唯一id，汇金outBizId
	TradeId int64 `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 汇金唯一id
	HjTradeNo int64 `json:"hj_trade_no,omitempty" xml:"hj_trade_no,omitempty"`
	// 汇金扣拥状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolParam = sync.Pool{
	New: func() any {
		return new(Param)
	},
}

// GetParam() 从对象池中获取Param
func GetParam() *Param {
	return poolParam.Get().(*Param)
}

// ReleaseParam 释放Param
func ReleaseParam(v *Param) {
	v.StoreCodes = v.StoreCodes[:0]
	v.Chainstores = v.Chainstores[:0]
	v.AppId = ""
	v.PartnerOrderCode = ""
	v.ChannelId = ""
	v.OrderCancelDesc = ""
	v.ServicePackageCode = ""
	v.MerchantCode = ""
	v.ChainstoreCode = ""
	v.Address = ""
	v.ChainStoreName = ""
	v.ChainStoreCode = ""
	v.ContactPhone = ""
	v.Longitude = ""
	v.Latitude = ""
	v.ErrorMsg = ""
	v.OrderCancelCode = 0
	v.OrderCancelReasonCode = 0
	v.PositionSource = 0
	v.TradeId = 0
	v.HjTradeNo = 0
	v.Success = false
	poolParam.Put(v)
}
