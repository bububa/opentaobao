package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappEleuserPhoneGetAPIRequest 获取饿了么用户信息 API请求
// taobao.miniapp.eleuser.phone.get
//
// 获取饿了么用户信息
type TaobaoMiniappEleuserPhoneGetAPIRequest struct {
	model.Params
}

// NewTaobaoMiniappEleuserPhoneGetRequest 初始化TaobaoMiniappEleuserPhoneGetAPIRequest对象
func NewTaobaoMiniappEleuserPhoneGetRequest() *TaobaoMiniappEleuserPhoneGetAPIRequest {
	return &TaobaoMiniappEleuserPhoneGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappEleuserPhoneGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappEleuserPhoneGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.eleuser.phone.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappEleuserPhoneGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappEleuserPhoneGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoMiniappEleuserPhoneGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappEleuserPhoneGetRequest()
	},
}

// GetTaobaoMiniappEleuserPhoneGetRequest 从 sync.Pool 获取 TaobaoMiniappEleuserPhoneGetAPIRequest
func GetTaobaoMiniappEleuserPhoneGetAPIRequest() *TaobaoMiniappEleuserPhoneGetAPIRequest {
	return poolTaobaoMiniappEleuserPhoneGetAPIRequest.Get().(*TaobaoMiniappEleuserPhoneGetAPIRequest)
}

// ReleaseTaobaoMiniappEleuserPhoneGetAPIRequest 将 TaobaoMiniappEleuserPhoneGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappEleuserPhoneGetAPIRequest(v *TaobaoMiniappEleuserPhoneGetAPIRequest) {
	v.Reset()
	poolTaobaoMiniappEleuserPhoneGetAPIRequest.Put(v)
}
