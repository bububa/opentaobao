package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardQryAPIRequest 查询卡实例 API请求
// alibaba.alsc.crm.card.qry
//
// 查询卡实例（优先使用卡实例ID查询，没有则用物理卡号查询）
type AlibabaAlscCrmCardQryAPIRequest struct {
	model.Params
	// 请求对象
	_paramQueryCardOpenReq *QueryCardOpenReq
}

// NewAlibabaAlscCrmCardQryRequest 初始化AlibabaAlscCrmCardQryAPIRequest对象
func NewAlibabaAlscCrmCardQryRequest() *AlibabaAlscCrmCardQryAPIRequest {
	return &AlibabaAlscCrmCardQryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardQryAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.qry"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardQryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamQueryCardOpenReq is ParamQueryCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardQryAPIRequest) SetParamQueryCardOpenReq(_paramQueryCardOpenReq *QueryCardOpenReq) error {
	r._paramQueryCardOpenReq = _paramQueryCardOpenReq
	r.Set("param_query_card_open_req", _paramQueryCardOpenReq)
	return nil
}

// GetParamQueryCardOpenReq ParamQueryCardOpenReq Getter
func (r AlibabaAlscCrmCardQryAPIRequest) GetParamQueryCardOpenReq() *QueryCardOpenReq {
	return r._paramQueryCardOpenReq
}
