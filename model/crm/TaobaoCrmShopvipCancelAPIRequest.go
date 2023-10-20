package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmshopvipcancelAPIRequest 卖家取消店铺vip的优惠 API请求
// taobao.crm.shopvip.cancel
//
// 此接口用于取消VIP优惠
type TaobaocrmshopvipcancelAPIRequest struct {
	model.Params
}

// NewTaobaocrmshopvipcancelRequest 初始化TaobaocrmshopvipcancelAPIRequest对象
func NewTaobaocrmshopvipcancelRequest() *TaobaocrmshopvipcancelAPIRequest {
	return &TaobaocrmshopvipcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmshopvipcancelAPIRequest) GetApiMethodName() string {
	return "taobao.crm.shopvip.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmshopvipcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmshopvipcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}
