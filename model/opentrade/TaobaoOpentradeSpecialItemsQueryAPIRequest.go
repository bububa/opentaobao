package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradespecialitemsqueryAPIRequest 专属下单获取商品绑定信息 API请求
// taobao.opentrade.special.items.query
//
// 专属下单获取商品绑定信息
type TaobaoopentradespecialitemsqueryAPIRequest struct {
	model.Params
	// 绑定专属下单场景的C端小程序ID
	_miniappId int64
}

// NewTaobaoopentradespecialitemsqueryRequest 初始化TaobaoopentradespecialitemsqueryAPIRequest对象
func NewTaobaoopentradespecialitemsqueryRequest() *TaobaoopentradespecialitemsqueryAPIRequest {
	return &TaobaoopentradespecialitemsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradespecialitemsqueryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.special.items.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradespecialitemsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradespecialitemsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniappId is MiniappId Setter
// 绑定专属下单场景的C端小程序ID
func (r *TaobaoopentradespecialitemsqueryAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// GetMiniappId MiniappId Getter
func (r TaobaoopentradespecialitemsqueryAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}
