package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappUserPhoneGetAPIRequest 获取当前授权用户手机号码 API请求
// taobao.miniapp.user.phone.get
//
// 在商家应用中，获取当前授权用户手机号码
type TaobaoMiniappUserPhoneGetAPIRequest struct {
	model.Params
}

// NewTaobaoMiniappUserPhoneGetRequest 初始化TaobaoMiniappUserPhoneGetAPIRequest对象
func NewTaobaoMiniappUserPhoneGetRequest() *TaobaoMiniappUserPhoneGetAPIRequest {
	return &TaobaoMiniappUserPhoneGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappUserPhoneGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappUserPhoneGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.user.phone.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappUserPhoneGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappUserPhoneGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoMiniappUserPhoneGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappUserPhoneGetRequest()
	},
}

// GetTaobaoMiniappUserPhoneGetRequest 从 sync.Pool 获取 TaobaoMiniappUserPhoneGetAPIRequest
func GetTaobaoMiniappUserPhoneGetAPIRequest() *TaobaoMiniappUserPhoneGetAPIRequest {
	return poolTaobaoMiniappUserPhoneGetAPIRequest.Get().(*TaobaoMiniappUserPhoneGetAPIRequest)
}

// ReleaseTaobaoMiniappUserPhoneGetAPIRequest 将 TaobaoMiniappUserPhoneGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappUserPhoneGetAPIRequest(v *TaobaoMiniappUserPhoneGetAPIRequest) {
	v.Reset()
	poolTaobaoMiniappUserPhoneGetAPIRequest.Put(v)
}
