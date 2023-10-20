package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAftersaleGetAPIRequest 查询用户售后服务模板 API请求
// taobao.aftersale.get
//
// 查询用户设置的售后服务模板，仅返回标题和id
type TaobaoAftersaleGetAPIRequest struct {
	model.Params
}

// NewTaobaoAftersaleGetRequest 初始化TaobaoAftersaleGetAPIRequest对象
func NewTaobaoAftersaleGetRequest() *TaobaoAftersaleGetAPIRequest {
	return &TaobaoAftersaleGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAftersaleGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAftersaleGetAPIRequest) GetApiMethodName() string {
	return "taobao.aftersale.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAftersaleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAftersaleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoAftersaleGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAftersaleGetRequest()
	},
}

// GetTaobaoAftersaleGetRequest 从 sync.Pool 获取 TaobaoAftersaleGetAPIRequest
func GetTaobaoAftersaleGetAPIRequest() *TaobaoAftersaleGetAPIRequest {
	return poolTaobaoAftersaleGetAPIRequest.Get().(*TaobaoAftersaleGetAPIRequest)
}

// ReleaseTaobaoAftersaleGetAPIRequest 将 TaobaoAftersaleGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAftersaleGetAPIRequest(v *TaobaoAftersaleGetAPIRequest) {
	v.Reset()
	poolTaobaoAftersaleGetAPIRequest.Put(v)
}
