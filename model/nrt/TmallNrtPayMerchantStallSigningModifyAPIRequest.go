package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtpaymerchantstallsigningmodifyAPIRequest 三级商户进件修改 API请求
// tmall.nrt.pay.merchant.stall.signing.modify
//
// 三级商户进件修改
type TmallnrtpaymerchantstallsigningmodifyAPIRequest struct {
	model.Params
	// 请求参数
	_req *StallSigningReqDto
}

// NewTmallnrtpaymerchantstallsigningmodifyRequest 初始化TmallnrtpaymerchantstallsigningmodifyAPIRequest对象
func NewTmallnrtpaymerchantstallsigningmodifyRequest() *TmallnrtpaymerchantstallsigningmodifyAPIRequest {
	return &TmallnrtpaymerchantstallsigningmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtpaymerchantstallsigningmodifyAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.pay.merchant.stall.signing.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtpaymerchantstallsigningmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtpaymerchantstallsigningmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 请求参数
func (r *TmallnrtpaymerchantstallsigningmodifyAPIRequest) SetReq(_req *StallSigningReqDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TmallnrtpaymerchantstallsigningmodifyAPIRequest) GetReq() *StallSigningReqDto {
	return r._req
}
