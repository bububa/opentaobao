package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoidlerecyclerefundcancleapplyAPIRequest 闲鱼回收取消退款申请V2 API请求
// taobao.idle.recycle.refund.cancleapply
//
// 回收商的回收订单取消退款申请
type TaobaoidlerecyclerefundcancleapplyAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewTaobaoidlerecyclerefundcancleapplyRequest 初始化TaobaoidlerecyclerefundcancleapplyAPIRequest对象
func NewTaobaoidlerecyclerefundcancleapplyRequest() *TaobaoidlerecyclerefundcancleapplyAPIRequest {
	return &TaobaoidlerecyclerefundcancleapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoidlerecyclerefundcancleapplyAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.cancleapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoidlerecyclerefundcancleapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoidlerecyclerefundcancleapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *TaobaoidlerecyclerefundcancleapplyAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoidlerecyclerefundcancleapplyAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
