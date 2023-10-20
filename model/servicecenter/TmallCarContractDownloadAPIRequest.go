package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarcontractdownloadAPIRequest 合同下载 API请求
// tmall.car.contract.download
//
// 目前天猫开新车会在线上签署一份合同，协议，需要一个个在已卖出打开，另存为pdf，人工一个个下载比较麻烦，期望通过接口直接读取pdf；
// 因为比较耗时，建议一个个下载，假设并发下载，很可能限流，每天的调用量有限；
type TmallcarcontractdownloadAPIRequest struct {
	model.Params
	// 天猫订单号
	_orderId int64
	// 是否下载html，true是html，false是pdf， html速度会快一点
	_html bool
}

// NewTmallcarcontractdownloadRequest 初始化TmallcarcontractdownloadAPIRequest对象
func NewTmallcarcontractdownloadRequest() *TmallcarcontractdownloadAPIRequest {
	return &TmallcarcontractdownloadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarcontractdownloadAPIRequest) GetApiMethodName() string {
	return "tmall.car.contract.download"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarcontractdownloadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarcontractdownloadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 天猫订单号
func (r *TmallcarcontractdownloadAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallcarcontractdownloadAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetHtml is Html Setter
// 是否下载html，true是html，false是pdf， html速度会快一点
func (r *TmallcarcontractdownloadAPIRequest) SetHtml(_html bool) error {
	r._html = _html
	r.Set("html", _html)
	return nil
}

// GetHtml Html Getter
func (r TmallcarcontractdownloadAPIRequest) GetHtml() bool {
	return r._html
}
