package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsonlinecancelAPIRequest 取消物流订单接口 API请求
// taobao.logistics.online.cancel
//
// 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
type TaobaologisticsonlinecancelAPIRequest struct {
	model.Params
	// 淘宝交易ID
	_tid int64
}

// NewTaobaologisticsonlinecancelRequest 初始化TaobaologisticsonlinecancelAPIRequest对象
func NewTaobaologisticsonlinecancelRequest() *TaobaologisticsonlinecancelAPIRequest {
	return &TaobaologisticsonlinecancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsonlinecancelAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.online.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsonlinecancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsonlinecancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 淘宝交易ID
func (r *TaobaologisticsonlinecancelAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaologisticsonlinecancelAPIRequest) GetTid() int64 {
	return r._tid
}
