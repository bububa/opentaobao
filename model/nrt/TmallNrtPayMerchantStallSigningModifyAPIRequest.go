package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtPayMerchantStallSigningModifyAPIRequest 三级商户进件修改 API请求
// tmall.nrt.pay.merchant.stall.signing.modify
//
// 三级商户进件修改
type TmallNrtPayMerchantStallSigningModifyAPIRequest struct {
	model.Params
	// 请求参数
	_req *StallSigningReqDto
}

// NewTmallNrtPayMerchantStallSigningModifyRequest 初始化TmallNrtPayMerchantStallSigningModifyAPIRequest对象
func NewTmallNrtPayMerchantStallSigningModifyRequest() *TmallNrtPayMerchantStallSigningModifyAPIRequest {
	return &TmallNrtPayMerchantStallSigningModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtPayMerchantStallSigningModifyAPIRequest) Reset() {
	r._req = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtPayMerchantStallSigningModifyAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.pay.merchant.stall.signing.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtPayMerchantStallSigningModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtPayMerchantStallSigningModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 请求参数
func (r *TmallNrtPayMerchantStallSigningModifyAPIRequest) SetReq(_req *StallSigningReqDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TmallNrtPayMerchantStallSigningModifyAPIRequest) GetReq() *StallSigningReqDto {
	return r._req
}

var poolTmallNrtPayMerchantStallSigningModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtPayMerchantStallSigningModifyRequest()
	},
}

// GetTmallNrtPayMerchantStallSigningModifyRequest 从 sync.Pool 获取 TmallNrtPayMerchantStallSigningModifyAPIRequest
func GetTmallNrtPayMerchantStallSigningModifyAPIRequest() *TmallNrtPayMerchantStallSigningModifyAPIRequest {
	return poolTmallNrtPayMerchantStallSigningModifyAPIRequest.Get().(*TmallNrtPayMerchantStallSigningModifyAPIRequest)
}

// ReleaseTmallNrtPayMerchantStallSigningModifyAPIRequest 将 TmallNrtPayMerchantStallSigningModifyAPIRequest 放入 sync.Pool
func ReleaseTmallNrtPayMerchantStallSigningModifyAPIRequest(v *TmallNrtPayMerchantStallSigningModifyAPIRequest) {
	v.Reset()
	poolTmallNrtPayMerchantStallSigningModifyAPIRequest.Put(v)
}
