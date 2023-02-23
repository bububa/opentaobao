package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointChkpntbypayAPIRequest 校验支付链路中的积分抵扣是否合法 API请求
// alibaba.alsc.crm.point.chkpntbypay
//
// 校验支付链路中的积分抵扣是否合法
type AlibabaAlscCrmPointChkpntbypayAPIRequest struct {
	model.Params
	// 入参
	_paramConsumePointByPayOpenReq *ConsumePointByPayOpenReq
}

// NewAlibabaAlscCrmPointChkpntbypayRequest 初始化AlibabaAlscCrmPointChkpntbypayAPIRequest对象
func NewAlibabaAlscCrmPointChkpntbypayRequest() *AlibabaAlscCrmPointChkpntbypayAPIRequest {
	return &AlibabaAlscCrmPointChkpntbypayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointChkpntbypayAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.chkpntbypay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointChkpntbypayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmPointChkpntbypayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamConsumePointByPayOpenReq is ParamConsumePointByPayOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointChkpntbypayAPIRequest) SetParamConsumePointByPayOpenReq(_paramConsumePointByPayOpenReq *ConsumePointByPayOpenReq) error {
	r._paramConsumePointByPayOpenReq = _paramConsumePointByPayOpenReq
	r.Set("param_consume_point_by_pay_open_req", _paramConsumePointByPayOpenReq)
	return nil
}

// GetParamConsumePointByPayOpenReq ParamConsumePointByPayOpenReq Getter
func (r AlibabaAlscCrmPointChkpntbypayAPIRequest) GetParamConsumePointByPayOpenReq() *ConsumePointByPayOpenReq {
	return r._paramConsumePointByPayOpenReq
}
