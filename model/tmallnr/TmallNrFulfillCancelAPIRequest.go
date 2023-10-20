package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrfulfillcancelAPIRequest 取消上门揽件 API请求
// tmall.nr.fulfill.cancel
//
// 新零售门店业务取消上门揽件
type TmallnrfulfillcancelAPIRequest struct {
	model.Params
	// 入参
	_req *NrCancelFulfillReqDto
}

// NewTmallnrfulfillcancelRequest 初始化TmallnrfulfillcancelAPIRequest对象
func NewTmallnrfulfillcancelRequest() *TmallnrfulfillcancelAPIRequest {
	return &TmallnrfulfillcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrfulfillcancelAPIRequest) GetApiMethodName() string {
	return "tmall.nr.fulfill.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrfulfillcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrfulfillcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 入参
func (r *TmallnrfulfillcancelAPIRequest) SetReq(_req *NrCancelFulfillReqDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TmallnrfulfillcancelAPIRequest) GetReq() *NrCancelFulfillReqDto {
	return r._req
}
