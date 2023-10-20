package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappEleuserinfoGetAPIRequest 获取饿了么用户信息 API请求
// taobao.miniapp.eleuserinfo.get
//
// 获取饿了么用户信息
type TaobaoMiniappEleuserinfoGetAPIRequest struct {
	model.Params
	// 怪兽业务方
	_bizScence string
}

// NewTaobaoMiniappEleuserinfoGetRequest 初始化TaobaoMiniappEleuserinfoGetAPIRequest对象
func NewTaobaoMiniappEleuserinfoGetRequest() *TaobaoMiniappEleuserinfoGetAPIRequest {
	return &TaobaoMiniappEleuserinfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappEleuserinfoGetAPIRequest) Reset() {
	r._bizScence = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappEleuserinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.eleuserinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappEleuserinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappEleuserinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizScence is BizScence Setter
// 怪兽业务方
func (r *TaobaoMiniappEleuserinfoGetAPIRequest) SetBizScence(_bizScence string) error {
	r._bizScence = _bizScence
	r.Set("biz_scence", _bizScence)
	return nil
}

// GetBizScence BizScence Getter
func (r TaobaoMiniappEleuserinfoGetAPIRequest) GetBizScence() string {
	return r._bizScence
}

var poolTaobaoMiniappEleuserinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappEleuserinfoGetRequest()
	},
}

// GetTaobaoMiniappEleuserinfoGetRequest 从 sync.Pool 获取 TaobaoMiniappEleuserinfoGetAPIRequest
func GetTaobaoMiniappEleuserinfoGetAPIRequest() *TaobaoMiniappEleuserinfoGetAPIRequest {
	return poolTaobaoMiniappEleuserinfoGetAPIRequest.Get().(*TaobaoMiniappEleuserinfoGetAPIRequest)
}

// ReleaseTaobaoMiniappEleuserinfoGetAPIRequest 将 TaobaoMiniappEleuserinfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappEleuserinfoGetAPIRequest(v *TaobaoMiniappEleuserinfoGetAPIRequest) {
	v.Reset()
	poolTaobaoMiniappEleuserinfoGetAPIRequest.Put(v)
}
