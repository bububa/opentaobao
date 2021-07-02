package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardQryphysicalAPIRequest 查询物理卡 API请求
// alibaba.alsc.crm.card.qryphysical
//
// 查询物理卡
type AlibabaAlscCrmCardQryphysicalAPIRequest struct {
	model.Params
	// 入参
	_paramQueryPhyCardOpenReq *QueryPhyCardOpenReq
}

// NewAlibabaAlscCrmCardQryphysicalRequest 初始化AlibabaAlscCrmCardQryphysicalAPIRequest对象
func NewAlibabaAlscCrmCardQryphysicalRequest() *AlibabaAlscCrmCardQryphysicalAPIRequest {
	return &AlibabaAlscCrmCardQryphysicalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardQryphysicalAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.qryphysical"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardQryphysicalAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamQueryPhyCardOpenReq Setter
// 入参
func (r *AlibabaAlscCrmCardQryphysicalAPIRequest) SetParamQueryPhyCardOpenReq(_paramQueryPhyCardOpenReq *QueryPhyCardOpenReq) error {
	r._paramQueryPhyCardOpenReq = _paramQueryPhyCardOpenReq
	r.Set("param_query_phy_card_open_req", _paramQueryPhyCardOpenReq)
	return nil
}

// Get ParamQueryPhyCardOpenReq Getter
func (r AlibabaAlscCrmCardQryphysicalAPIRequest) GetParamQueryPhyCardOpenReq() *QueryPhyCardOpenReq {
	return r._paramQueryPhyCardOpenReq
}
