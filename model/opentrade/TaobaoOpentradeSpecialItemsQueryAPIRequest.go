package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialItemsQueryAPIRequest 专属下单获取商品绑定信息 API请求
// taobao.opentrade.special.items.query
//
// 专属下单获取商品绑定信息
type TaobaoOpentradeSpecialItemsQueryAPIRequest struct {
	model.Params
	// 绑定专属下单场景的C端小程序ID
	_miniappId int64
}

// NewTaobaoOpentradeSpecialItemsQueryRequest 初始化TaobaoOpentradeSpecialItemsQueryAPIRequest对象
func NewTaobaoOpentradeSpecialItemsQueryRequest() *TaobaoOpentradeSpecialItemsQueryAPIRequest {
	return &TaobaoOpentradeSpecialItemsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialItemsQueryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.special.items.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialItemsQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMiniappId is MiniappId Setter
// 绑定专属下单场景的C端小程序ID
func (r *TaobaoOpentradeSpecialItemsQueryAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// GetMiniappId MiniappId Getter
func (r TaobaoOpentradeSpecialItemsQueryAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}
