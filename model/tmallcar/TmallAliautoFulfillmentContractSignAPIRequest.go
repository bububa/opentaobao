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
	_req *SignContractReq
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
func (r TmallAliautoFulfillmentContractSignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoFulfillmentContractSignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 入参
func (r *TmallAliautoFulfillmentContractSignAPIRequest) SetReq(_req *SignContractReq) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TmallAliautoFulfillmentContractSignAPIRequest) GetReq() *SignContractReq {
	return r._req
}
