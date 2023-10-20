package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardqryAPIRequest 查询卡实例 API请求
// alibaba.alsc.crm.card.qry
//
// 查询卡实例（优先使用卡实例ID查询，没有则用物理卡号查询）
type AlibabaalsccrmcardqryAPIRequest struct {
	model.Params
	// 请求对象
	_paramQueryCardOpenReq *QueryCardOpenReq
}

// NewAlibabaalsccrmcardqryRequest 初始化AlibabaalsccrmcardqryAPIRequest对象
func NewAlibabaalsccrmcardqryRequest() *AlibabaalsccrmcardqryAPIRequest {
	return &AlibabaalsccrmcardqryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardqryAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.qry"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardqryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardqryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryCardOpenReq is ParamQueryCardOpenReq Setter
// 请求对象
func (r *AlibabaalsccrmcardqryAPIRequest) SetParamQueryCardOpenReq(_paramQueryCardOpenReq *QueryCardOpenReq) error {
	r._paramQueryCardOpenReq = _paramQueryCardOpenReq
	r.Set("param_query_card_open_req", _paramQueryCardOpenReq)
	return nil
}

// GetParamQueryCardOpenReq ParamQueryCardOpenReq Getter
func (r AlibabaalsccrmcardqryAPIRequest) GetParamQueryCardOpenReq() *QueryCardOpenReq {
	return r._paramQueryCardOpenReq
}
