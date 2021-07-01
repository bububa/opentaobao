package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrFulfillCancelAPIRequest
取消上门揽件 API请求
tmall.nr.fulfill.cancel

新零售门店业务取消上门揽件 */
type TmallNrFulfillCancelAPIRequest struct {
	model.Params
	// 入参
	_req *NrCancelFulfillReqDto
}

// NewTmallNrFulfillCancelRequest 初始化TmallNrFulfillCancelAPIRequest对象
func NewTmallNrFulfillCancelRequest() *TmallNrFulfillCancelAPIRequest {
	return &TmallNrFulfillCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrFulfillCancelAPIRequest) GetApiMethodName() string {
	return "tmall.nr.fulfill.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrFulfillCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Req Setter
// 入参
func (r *TmallNrFulfillCancelAPIRequest) SetReq(_req *NrCancelFulfillReqDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// Get Req Getter
func (r TmallNrFulfillCancelAPIRequest) GetReq() *NrCancelFulfillReqDto {
	return r._req
}
