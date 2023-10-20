package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarCarefreeDetailGetAPIRequest 查询业务单信息 API请求
// tmall.car.carefree.detail.get
//
// 查询业务单信息
type TmallCarCarefreeDetailGetAPIRequest struct {
	model.Params
	// 查询参数对象
	_param0 *CarefreeDetailQueryReq
}

// NewTmallCarCarefreeDetailGetRequest 初始化TmallCarCarefreeDetailGetAPIRequest对象
func NewTmallCarCarefreeDetailGetRequest() *TmallCarCarefreeDetailGetAPIRequest {
	return &TmallCarCarefreeDetailGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarCarefreeDetailGetAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarCarefreeDetailGetAPIRequest) GetApiMethodName() string {
	return "tmall.car.carefree.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarCarefreeDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarCarefreeDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询参数对象
func (r *TmallCarCarefreeDetailGetAPIRequest) SetParam0(_param0 *CarefreeDetailQueryReq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallCarCarefreeDetailGetAPIRequest) GetParam0() *CarefreeDetailQueryReq {
	return r._param0
}

var poolTmallCarCarefreeDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarCarefreeDetailGetRequest()
	},
}

// GetTmallCarCarefreeDetailGetRequest 从 sync.Pool 获取 TmallCarCarefreeDetailGetAPIRequest
func GetTmallCarCarefreeDetailGetAPIRequest() *TmallCarCarefreeDetailGetAPIRequest {
	return poolTmallCarCarefreeDetailGetAPIRequest.Get().(*TmallCarCarefreeDetailGetAPIRequest)
}

// ReleaseTmallCarCarefreeDetailGetAPIRequest 将 TmallCarCarefreeDetailGetAPIRequest 放入 sync.Pool
func ReleaseTmallCarCarefreeDetailGetAPIRequest(v *TmallCarCarefreeDetailGetAPIRequest) {
	v.Reset()
	poolTmallCarCarefreeDetailGetAPIRequest.Put(v)
}
