package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillCancelAPIRequest 取消上门揽件 API请求
// tmall.nr.fulfill.cancel
//
// 新零售门店业务取消上门揽件
type TmallNrFulfillCancelAPIRequest struct {
	model.Params
	// 入参
	_req *NrCancelFulfillReqDto
}

// NewTmallNrFulfillCancelRequest 初始化TmallNrFulfillCancelAPIRequest对象
func NewTmallNrFulfillCancelRequest() *TmallNrFulfillCancelAPIRequest {
	return &TmallNrFulfillCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrFulfillCancelAPIRequest) Reset() {
	r._req = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrFulfillCancelAPIRequest) GetApiMethodName() string {
	return "tmall.nr.fulfill.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrFulfillCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrFulfillCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 入参
func (r *TmallNrFulfillCancelAPIRequest) SetReq(_req *NrCancelFulfillReqDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TmallNrFulfillCancelAPIRequest) GetReq() *NrCancelFulfillReqDto {
	return r._req
}

var poolTmallNrFulfillCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrFulfillCancelRequest()
	},
}

// GetTmallNrFulfillCancelRequest 从 sync.Pool 获取 TmallNrFulfillCancelAPIRequest
func GetTmallNrFulfillCancelAPIRequest() *TmallNrFulfillCancelAPIRequest {
	return poolTmallNrFulfillCancelAPIRequest.Get().(*TmallNrFulfillCancelAPIRequest)
}

// ReleaseTmallNrFulfillCancelAPIRequest 将 TmallNrFulfillCancelAPIRequest 放入 sync.Pool
func ReleaseTmallNrFulfillCancelAPIRequest(v *TmallNrFulfillCancelAPIRequest) {
	v.Reset()
	poolTmallNrFulfillCancelAPIRequest.Put(v)
}
