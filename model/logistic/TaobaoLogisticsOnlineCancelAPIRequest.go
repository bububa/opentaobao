package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOnlineCancelAPIRequest 取消物流订单接口 API请求
// taobao.logistics.online.cancel
//
// 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
type TaobaoLogisticsOnlineCancelAPIRequest struct {
	model.Params
	// 淘宝交易ID
	_tid int64
}

// NewTaobaoLogisticsOnlineCancelRequest 初始化TaobaoLogisticsOnlineCancelAPIRequest对象
func NewTaobaoLogisticsOnlineCancelRequest() *TaobaoLogisticsOnlineCancelAPIRequest {
	return &TaobaoLogisticsOnlineCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOnlineCancelAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.online.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOnlineCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsOnlineCancelAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoLogisticsOnlineCancelAPIRequest) GetTid() int64 {
	return r._tid
}
