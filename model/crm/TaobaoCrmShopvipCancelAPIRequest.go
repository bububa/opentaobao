package crm

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmShopvipCancelAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmShopvipCancelAPIRequest) GetApiMethodName() string {
	return "taobao.crm.shopvip.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmShopvipCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmShopvipCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoCrmShopvipCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmShopvipCancelRequest()
	},
}

// GetTaobaoCrmShopvipCancelRequest 从 sync.Pool 获取 TaobaoCrmShopvipCancelAPIRequest
func GetTaobaoCrmShopvipCancelAPIRequest() *TaobaoCrmShopvipCancelAPIRequest {
	return poolTaobaoCrmShopvipCancelAPIRequest.Get().(*TaobaoCrmShopvipCancelAPIRequest)
}

// ReleaseTaobaoCrmShopvipCancelAPIRequest 将 TaobaoCrmShopvipCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmShopvipCancelAPIRequest(v *TaobaoCrmShopvipCancelAPIRequest) {
	v.Reset()
	poolTaobaoCrmShopvipCancelAPIRequest.Put(v)
}
