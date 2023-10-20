package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappUserInfoGetAPIRequest 用户开放信息获取 API请求
// taobao.miniapp.userInfo.get
//
// 获取用户的 openId，snsNick（如果用户设置过的话），和加密头像链接
type TaobaoMiniappUserInfoGetAPIRequest struct {
	model.Params
}

// NewTaobaoMiniappUserInfoGetRequest 初始化TaobaoMiniappUserInfoGetAPIRequest对象
func NewTaobaoMiniappUserInfoGetRequest() *TaobaoMiniappUserInfoGetAPIRequest {
	return &TaobaoMiniappUserInfoGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappUserInfoGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappUserInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.userInfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappUserInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappUserInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoMiniappUserInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappUserInfoGetRequest()
	},
}

// GetTaobaoMiniappUserInfoGetRequest 从 sync.Pool 获取 TaobaoMiniappUserInfoGetAPIRequest
func GetTaobaoMiniappUserInfoGetAPIRequest() *TaobaoMiniappUserInfoGetAPIRequest {
	return poolTaobaoMiniappUserInfoGetAPIRequest.Get().(*TaobaoMiniappUserInfoGetAPIRequest)
}

// ReleaseTaobaoMiniappUserInfoGetAPIRequest 将 TaobaoMiniappUserInfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappUserInfoGetAPIRequest(v *TaobaoMiniappUserInfoGetAPIRequest) {
	v.Reset()
	poolTaobaoMiniappUserInfoGetAPIRequest.Put(v)
}
