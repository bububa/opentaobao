package tblogistics

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsOnlineCancelAPIRequest) Reset() {
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOnlineCancelAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.online.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOnlineCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsOnlineCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsOnlineCancelAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoLogisticsOnlineCancelAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoLogisticsOnlineCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsOnlineCancelRequest()
	},
}

// GetTaobaoLogisticsOnlineCancelRequest 从 sync.Pool 获取 TaobaoLogisticsOnlineCancelAPIRequest
func GetTaobaoLogisticsOnlineCancelAPIRequest() *TaobaoLogisticsOnlineCancelAPIRequest {
	return poolTaobaoLogisticsOnlineCancelAPIRequest.Get().(*TaobaoLogisticsOnlineCancelAPIRequest)
}

// ReleaseTaobaoLogisticsOnlineCancelAPIRequest 将 TaobaoLogisticsOnlineCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsOnlineCancelAPIRequest(v *TaobaoLogisticsOnlineCancelAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsOnlineCancelAPIRequest.Put(v)
}
