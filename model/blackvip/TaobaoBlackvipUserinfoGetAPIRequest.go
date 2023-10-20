package blackvip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBlackvipUserinfoGetAPIRequest 88VIP用户信息查询 API请求
// taobao.blackvip.userinfo.get
//
// 查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
type TaobaoBlackvipUserinfoGetAPIRequest struct {
	model.Params
}

// NewTaobaoBlackvipUserinfoGetRequest 初始化TaobaoBlackvipUserinfoGetAPIRequest对象
func NewTaobaoBlackvipUserinfoGetRequest() *TaobaoBlackvipUserinfoGetAPIRequest {
	return &TaobaoBlackvipUserinfoGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBlackvipUserinfoGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBlackvipUserinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.blackvip.userinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBlackvipUserinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBlackvipUserinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoBlackvipUserinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBlackvipUserinfoGetRequest()
	},
}

// GetTaobaoBlackvipUserinfoGetRequest 从 sync.Pool 获取 TaobaoBlackvipUserinfoGetAPIRequest
func GetTaobaoBlackvipUserinfoGetAPIRequest() *TaobaoBlackvipUserinfoGetAPIRequest {
	return poolTaobaoBlackvipUserinfoGetAPIRequest.Get().(*TaobaoBlackvipUserinfoGetAPIRequest)
}

// ReleaseTaobaoBlackvipUserinfoGetAPIRequest 将 TaobaoBlackvipUserinfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBlackvipUserinfoGetAPIRequest(v *TaobaoBlackvipUserinfoGetAPIRequest) {
	v.Reset()
	poolTaobaoBlackvipUserinfoGetAPIRequest.Put(v)
}
