package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemHscodeDetailGetAPIRequest 通过hscode获取计量单位 API请求
// tmall.item.hscode.detail.get
//
// 通过hscode获取计量单位和销售单位
type TmallItemHscodeDetailGetAPIRequest struct {
	model.Params
	// hscode
	_hscode string
}

// NewTmallItemHscodeDetailGetRequest 初始化TmallItemHscodeDetailGetAPIRequest对象
func NewTmallItemHscodeDetailGetRequest() *TmallItemHscodeDetailGetAPIRequest {
	return &TmallItemHscodeDetailGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemHscodeDetailGetAPIRequest) Reset() {
	r._hscode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemHscodeDetailGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.hscode.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemHscodeDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemHscodeDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHscode is Hscode Setter
// hscode
func (r *TmallItemHscodeDetailGetAPIRequest) SetHscode(_hscode string) error {
	r._hscode = _hscode
	r.Set("hscode", _hscode)
	return nil
}

// GetHscode Hscode Getter
func (r TmallItemHscodeDetailGetAPIRequest) GetHscode() string {
	return r._hscode
}

var poolTmallItemHscodeDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemHscodeDetailGetRequest()
	},
}

// GetTmallItemHscodeDetailGetRequest 从 sync.Pool 获取 TmallItemHscodeDetailGetAPIRequest
func GetTmallItemHscodeDetailGetAPIRequest() *TmallItemHscodeDetailGetAPIRequest {
	return poolTmallItemHscodeDetailGetAPIRequest.Get().(*TmallItemHscodeDetailGetAPIRequest)
}

// ReleaseTmallItemHscodeDetailGetAPIRequest 将 TmallItemHscodeDetailGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemHscodeDetailGetAPIRequest(v *TmallItemHscodeDetailGetAPIRequest) {
	v.Reset()
	poolTmallItemHscodeDetailGetAPIRequest.Put(v)
}
