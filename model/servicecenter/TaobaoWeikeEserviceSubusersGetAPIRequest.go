package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikeEserviceSubusersGetAPIRequest 客服外包订单分配的商家子账号列表 API请求
// taobao.weike.eservice.subusers.get
//
// 获取客服外包订单分配的商家子账号列表，以及授权状态
type TaobaoWeikeEserviceSubusersGetAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
}

// NewTaobaoWeikeEserviceSubusersGetRequest 初始化TaobaoWeikeEserviceSubusersGetAPIRequest对象
func NewTaobaoWeikeEserviceSubusersGetRequest() *TaobaoWeikeEserviceSubusersGetAPIRequest {
	return &TaobaoWeikeEserviceSubusersGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWeikeEserviceSubusersGetAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceSubusersGetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.eservice.subusers.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceSubusersGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWeikeEserviceSubusersGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaoWeikeEserviceSubusersGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoWeikeEserviceSubusersGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTaobaoWeikeEserviceSubusersGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWeikeEserviceSubusersGetRequest()
	},
}

// GetTaobaoWeikeEserviceSubusersGetRequest 从 sync.Pool 获取 TaobaoWeikeEserviceSubusersGetAPIRequest
func GetTaobaoWeikeEserviceSubusersGetAPIRequest() *TaobaoWeikeEserviceSubusersGetAPIRequest {
	return poolTaobaoWeikeEserviceSubusersGetAPIRequest.Get().(*TaobaoWeikeEserviceSubusersGetAPIRequest)
}

// ReleaseTaobaoWeikeEserviceSubusersGetAPIRequest 将 TaobaoWeikeEserviceSubusersGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWeikeEserviceSubusersGetAPIRequest(v *TaobaoWeikeEserviceSubusersGetAPIRequest) {
	v.Reset()
	poolTaobaoWeikeEserviceSubusersGetAPIRequest.Put(v)
}
