package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardqryphysicalAPIRequest 查询物理卡 API请求
// alibaba.alsc.crm.card.qryphysical
//
// 查询物理卡
type AlibabaalsccrmcardqryphysicalAPIRequest struct {
	model.Params
	// 入参
	_paramQueryPhyCardOpenReq *QueryPhyCardOpenReq
}

// NewAlibabaalsccrmcardqryphysicalRequest 初始化AlibabaalsccrmcardqryphysicalAPIRequest对象
func NewAlibabaalsccrmcardqryphysicalRequest() *AlibabaalsccrmcardqryphysicalAPIRequest {
	return &AlibabaalsccrmcardqryphysicalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardqryphysicalAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.qryphysical"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardqryphysicalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardqryphysicalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryPhyCardOpenReq is ParamQueryPhyCardOpenReq Setter
// 入参
func (r *AlibabaalsccrmcardqryphysicalAPIRequest) SetParamQueryPhyCardOpenReq(_paramQueryPhyCardOpenReq *QueryPhyCardOpenReq) error {
	r._paramQueryPhyCardOpenReq = _paramQueryPhyCardOpenReq
	r.Set("param_query_phy_card_open_req", _paramQueryPhyCardOpenReq)
	return nil
}

// GetParamQueryPhyCardOpenReq ParamQueryPhyCardOpenReq Getter
func (r AlibabaalsccrmcardqryphysicalAPIRequest) GetParamQueryPhyCardOpenReq() *QueryPhyCardOpenReq {
	return r._paramQueryPhyCardOpenReq
}
