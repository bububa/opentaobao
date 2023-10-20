package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOmniSaasOrderCreateAPIRequest 订单创建接口 API请求
// alibaba.omni.saas.order.create
//
// 服务商利用现有的saas系统和阿里完成交易系统的对接
type AlibabaOmniSaasOrderCreateAPIRequest struct {
	model.Params
	// 商品列表
	_goodsDetails []GoodsDetail
	// 商家自有优惠
	_couponInfos []CouponInfo
	// 买家标识，淘系用户或用户手机号。当支付渠道为支付宝时，此字段为淘宝会员码或支付宝付款码。(当前仅支持淘系用户，手机号下单稍后开放)
	_buyerId string
	// ALIPAY：支付宝用户；TAOBAO：淘宝会员码；MOBILE：手机号
	_buyerIdType string
	// 门店ID
	_storeId string
	// 收银设备类型
	_device string
	// 收银设备号
	_deviceNo string
	// 收银员标识
	_operatorId string
	// ALIPAY：支付宝付款；BANK_CARD：刷卡
	_payChannel string
	// PLACE：淘宝商户中心门店ID；CUSTOM：商户自有门店编码，需要维护到淘宝商户中心
	_storeIdType string
	// 请求号，用于标识一次请求
	_requestNo string
}

// NewAlibabaOmniSaasOrderCreateRequest 初始化AlibabaOmniSaasOrderCreateAPIRequest对象
func NewAlibabaOmniSaasOrderCreateRequest() *AlibabaOmniSaasOrderCreateAPIRequest {
	return &AlibabaOmniSaasOrderCreateAPIRequest{
		Params: model.NewParams(11),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOmniSaasOrderCreateAPIRequest) Reset() {
	r._goodsDetails = r._goodsDetails[:0]
	r._couponInfos = r._couponInfos[:0]
	r._buyerId = ""
	r._buyerIdType = ""
	r._storeId = ""
	r._device = ""
	r._deviceNo = ""
	r._operatorId = ""
	r._payChannel = ""
	r._storeIdType = ""
	r._requestNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.omni.saas.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsDetails is GoodsDetails Setter
// 商品列表
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetGoodsDetails(_goodsDetails []GoodsDetail) error {
	r._goodsDetails = _goodsDetails
	r.Set("goods_details", _goodsDetails)
	return nil
}

// GetGoodsDetails GoodsDetails Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetGoodsDetails() []GoodsDetail {
	return r._goodsDetails
}

// SetCouponInfos is CouponInfos Setter
// 商家自有优惠
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetCouponInfos(_couponInfos []CouponInfo) error {
	r._couponInfos = _couponInfos
	r.Set("coupon_infos", _couponInfos)
	return nil
}

// GetCouponInfos CouponInfos Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetCouponInfos() []CouponInfo {
	return r._couponInfos
}

// SetBuyerId is BuyerId Setter
// 买家标识，淘系用户或用户手机号。当支付渠道为支付宝时，此字段为淘宝会员码或支付宝付款码。(当前仅支持淘系用户，手机号下单稍后开放)
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetBuyerId(_buyerId string) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetBuyerId() string {
	return r._buyerId
}

// SetBuyerIdType is BuyerIdType Setter
// ALIPAY：支付宝用户；TAOBAO：淘宝会员码；MOBILE：手机号
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetBuyerIdType(_buyerIdType string) error {
	r._buyerIdType = _buyerIdType
	r.Set("buyer_id_type", _buyerIdType)
	return nil
}

// GetBuyerIdType BuyerIdType Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetBuyerIdType() string {
	return r._buyerIdType
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetDevice is Device Setter
// 收银设备类型
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetDevice(_device string) error {
	r._device = _device
	r.Set("device", _device)
	return nil
}

// GetDevice Device Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetDevice() string {
	return r._device
}

// SetDeviceNo is DeviceNo Setter
// 收银设备号
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetDeviceNo(_deviceNo string) error {
	r._deviceNo = _deviceNo
	r.Set("device_no", _deviceNo)
	return nil
}

// GetDeviceNo DeviceNo Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetDeviceNo() string {
	return r._deviceNo
}

// SetOperatorId is OperatorId Setter
// 收银员标识
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetOperatorId(_operatorId string) error {
	r._operatorId = _operatorId
	r.Set("operator_id", _operatorId)
	return nil
}

// GetOperatorId OperatorId Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetOperatorId() string {
	return r._operatorId
}

// SetPayChannel is PayChannel Setter
// ALIPAY：支付宝付款；BANK_CARD：刷卡
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetPayChannel(_payChannel string) error {
	r._payChannel = _payChannel
	r.Set("pay_channel", _payChannel)
	return nil
}

// GetPayChannel PayChannel Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetPayChannel() string {
	return r._payChannel
}

// SetStoreIdType is StoreIdType Setter
// PLACE：淘宝商户中心门店ID；CUSTOM：商户自有门店编码，需要维护到淘宝商户中心
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetStoreIdType(_storeIdType string) error {
	r._storeIdType = _storeIdType
	r.Set("store_id_type", _storeIdType)
	return nil
}

// GetStoreIdType StoreIdType Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetStoreIdType() string {
	return r._storeIdType
}

// SetRequestNo is RequestNo Setter
// 请求号，用于标识一次请求
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetRequestNo(_requestNo string) error {
	r._requestNo = _requestNo
	r.Set("request_no", _requestNo)
	return nil
}

// GetRequestNo RequestNo Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetRequestNo() string {
	return r._requestNo
}

var poolAlibabaOmniSaasOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOmniSaasOrderCreateRequest()
	},
}

// GetAlibabaOmniSaasOrderCreateRequest 从 sync.Pool 获取 AlibabaOmniSaasOrderCreateAPIRequest
func GetAlibabaOmniSaasOrderCreateAPIRequest() *AlibabaOmniSaasOrderCreateAPIRequest {
	return poolAlibabaOmniSaasOrderCreateAPIRequest.Get().(*AlibabaOmniSaasOrderCreateAPIRequest)
}

// ReleaseAlibabaOmniSaasOrderCreateAPIRequest 将 AlibabaOmniSaasOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaOmniSaasOrderCreateAPIRequest(v *AlibabaOmniSaasOrderCreateAPIRequest) {
	v.Reset()
	poolAlibabaOmniSaasOrderCreateAPIRequest.Put(v)
}
