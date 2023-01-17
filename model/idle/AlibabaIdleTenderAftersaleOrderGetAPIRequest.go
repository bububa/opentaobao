package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderAftersaleOrderGetAPIRequest 闲鱼帮卖售后服务单查询 API请求
// alibaba.idle.tender.aftersale.order.get
//
// 闲鱼帮卖售后服务单查询
type AlibabaIdleTenderAftersaleOrderGetAPIRequest struct {
	model.Params
	// 主订单号
	_mainOrderId int64
	// 服务申请单id
	_applyId int64
}

// NewAlibabaIdleTenderAftersaleOrderGetRequest 初始化AlibabaIdleTenderAftersaleOrderGetAPIRequest对象
func NewAlibabaIdleTenderAftersaleOrderGetRequest() *AlibabaIdleTenderAftersaleOrderGetAPIRequest {
	return &AlibabaIdleTenderAftersaleOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTenderAftersaleOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.aftersale.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTenderAftersaleOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleTenderAftersaleOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 主订单号
func (r *AlibabaIdleTenderAftersaleOrderGetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r AlibabaIdleTenderAftersaleOrderGetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetApplyId is ApplyId Setter
// 服务申请单id
func (r *AlibabaIdleTenderAftersaleOrderGetAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaIdleTenderAftersaleOrderGetAPIRequest) GetApplyId() int64 {
	return r._applyId
}
