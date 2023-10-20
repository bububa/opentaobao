package subuser

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubusersInfoQueryAPIRequest 根据当前子账号登陆态，获取该子账号基本信息 API请求
// taobao.subusers.info.query
//
// 根据当前子账号登陆态，获取该子账号基本信息
type TaobaoSubusersInfoQueryAPIRequest struct {
	model.Params
	// 站点信息，淘宝天猫传0，1688传3
	_site int64
}

// NewTaobaoSubusersInfoQueryRequest 初始化TaobaoSubusersInfoQueryAPIRequest对象
func NewTaobaoSubusersInfoQueryRequest() *TaobaoSubusersInfoQueryAPIRequest {
	return &TaobaoSubusersInfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubusersInfoQueryAPIRequest) Reset() {
	r._site = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubusersInfoQueryAPIRequest) GetApiMethodName() string {
	return "taobao.subusers.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubusersInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubusersInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSite is Site Setter
// 站点信息，淘宝天猫传0，1688传3
func (r *TaobaoSubusersInfoQueryAPIRequest) SetSite(_site int64) error {
	r._site = _site
	r.Set("site", _site)
	return nil
}

// GetSite Site Getter
func (r TaobaoSubusersInfoQueryAPIRequest) GetSite() int64 {
	return r._site
}

var poolTaobaoSubusersInfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubusersInfoQueryRequest()
	},
}

// GetTaobaoSubusersInfoQueryRequest 从 sync.Pool 获取 TaobaoSubusersInfoQueryAPIRequest
func GetTaobaoSubusersInfoQueryAPIRequest() *TaobaoSubusersInfoQueryAPIRequest {
	return poolTaobaoSubusersInfoQueryAPIRequest.Get().(*TaobaoSubusersInfoQueryAPIRequest)
}

// ReleaseTaobaoSubusersInfoQueryAPIRequest 将 TaobaoSubusersInfoQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubusersInfoQueryAPIRequest(v *TaobaoSubusersInfoQueryAPIRequest) {
	v.Reset()
	poolTaobaoSubusersInfoQueryAPIRequest.Put(v)
}
