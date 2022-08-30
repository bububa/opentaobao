package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoFulfillmentContractSignAPIRequest 合同签署 API请求
// tmall.aliauto.fulfillment.contract.sign
//
// 商家回传用户签署的合同信息
type TmallAliautoFulfillmentContractSignAPIRequest struct {
	model.Params
	// 入参
	_req *ConsignContractReq
}

// NewTmallAliautoFulfillmentContractSignRequest 初始化TmallAliautoFulfillmentContractSignAPIRequest对象
func NewTmallAliautoFulfillmentContractSignRequest() *TmallAliautoFulfillmentContractSignAPIRequest {
	return &TmallAliautoFulfillmentContractSignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoFulfillmentContractSignAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.fulfillment.contract.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoFulfillmentContractSignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReq is Req Setter
// 入参
func (r *TmallAliautoFulfillmentContractSignAPIRequest) SetReq(_req *ConsignContractReq) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TmallAliautoFulfillmentContractSignAPIRequest) GetReq() *ConsignContractReq {
	return r._req
}
