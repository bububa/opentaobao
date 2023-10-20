package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemOutercodeGetAPIRequest 根据outerCode查询商品 API请求
// taobao.scitem.outercode.get
//
// 根据outerCode查询商品
type TaobaoScitemOutercodeGetAPIRequest struct {
	model.Params
	// 商品编码
	_outerCode string
}

// NewTaobaoScitemOutercodeGetRequest 初始化TaobaoScitemOutercodeGetAPIRequest对象
func NewTaobaoScitemOutercodeGetRequest() *TaobaoScitemOutercodeGetAPIRequest {
	return &TaobaoScitemOutercodeGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoScitemOutercodeGetAPIRequest) Reset() {
	r._outerCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoScitemOutercodeGetAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.outercode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoScitemOutercodeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoScitemOutercodeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCode is OuterCode Setter
// 商品编码
func (r *TaobaoScitemOutercodeGetAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TaobaoScitemOutercodeGetAPIRequest) GetOuterCode() string {
	return r._outerCode
}

var poolTaobaoScitemOutercodeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoScitemOutercodeGetRequest()
	},
}

// GetTaobaoScitemOutercodeGetRequest 从 sync.Pool 获取 TaobaoScitemOutercodeGetAPIRequest
func GetTaobaoScitemOutercodeGetAPIRequest() *TaobaoScitemOutercodeGetAPIRequest {
	return poolTaobaoScitemOutercodeGetAPIRequest.Get().(*TaobaoScitemOutercodeGetAPIRequest)
}

// ReleaseTaobaoScitemOutercodeGetAPIRequest 将 TaobaoScitemOutercodeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoScitemOutercodeGetAPIRequest(v *TaobaoScitemOutercodeGetAPIRequest) {
	v.Reset()
	poolTaobaoScitemOutercodeGetAPIRequest.Put(v)
}
