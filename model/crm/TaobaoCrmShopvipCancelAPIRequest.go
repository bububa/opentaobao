package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmShopvipCancelAPIRequest 卖家取消店铺vip的优惠 API请求
// taobao.crm.shopvip.cancel
//
// 此接口用于取消VIP优惠
type TaobaoCrmShopvipCancelAPIRequest struct {
	model.Params
}

// NewTaobaoCrmShopvipCancelRequest 初始化TaobaoCrmShopvipCancelAPIRequest对象
func NewTaobaoCrmShopvipCancelRequest() *TaobaoCrmShopvipCancelAPIRequest {
	return &TaobaoCrmShopvipCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmShopvipCancelAPIRequest) GetApiMethodName() string {
	return "taobao.crm.shopvip.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmShopvipCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
