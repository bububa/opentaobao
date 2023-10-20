package aliexpresssumaitong

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresstradeorderopenqueryAPIRequest Aliexpress开放平台订单查询 API请求
// aliexpress.trade.order.open.query
//
// Aliexpress开放平台订单信息查询
type AliexpresstradeorderopenqueryAPIRequest struct {
	model.Params
	// 订单号
	_orderIds []int64
	// 外部订单号
	_outIds []string
	// 小程序appId
	_openAppKey string
	// 业务编码
	_bizCode string
	// 买家用户id
	_buyerId int64
}

// NewAliexpresstradeorderopenqueryRequest 初始化AliexpresstradeorderopenqueryAPIRequest对象
func NewAliexpresstradeorderopenqueryRequest() *AliexpresstradeorderopenqueryAPIRequest {
	return &AliexpresstradeorderopenqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresstradeorderopenqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.trade.order.open.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresstradeorderopenqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresstradeorderopenqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderIds is OrderIds Setter
// 订单号
func (r *AliexpresstradeorderopenqueryAPIRequest) SetOrderIds(_orderIds []int64) error {
	r._orderIds = _orderIds
	r.Set("order_ids", _orderIds)
	return nil
}

// GetOrderIds OrderIds Getter
func (r AliexpresstradeorderopenqueryAPIRequest) GetOrderIds() []int64 {
	return r._orderIds
}

// SetOutIds is OutIds Setter
// 外部订单号
func (r *AliexpresstradeorderopenqueryAPIRequest) SetOutIds(_outIds []string) error {
	r._outIds = _outIds
	r.Set("out_ids", _outIds)
	return nil
}

// GetOutIds OutIds Getter
func (r AliexpresstradeorderopenqueryAPIRequest) GetOutIds() []string {
	return r._outIds
}

// SetOpenAppKey is OpenAppKey Setter
// 小程序appId
func (r *AliexpresstradeorderopenqueryAPIRequest) SetOpenAppKey(_openAppKey string) error {
	r._openAppKey = _openAppKey
	r.Set("open_app_key", _openAppKey)
	return nil
}

// GetOpenAppKey OpenAppKey Getter
func (r AliexpresstradeorderopenqueryAPIRequest) GetOpenAppKey() string {
	return r._openAppKey
}

// SetBizCode is BizCode Setter
// 业务编码
func (r *AliexpresstradeorderopenqueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r AliexpresstradeorderopenqueryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetBuyerId is BuyerId Setter
// 买家用户id
func (r *AliexpresstradeorderopenqueryAPIRequest) SetBuyerId(_buyerId int64) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r AliexpresstradeorderopenqueryAPIRequest) GetBuyerId() int64 {
	return r._buyerId
}
