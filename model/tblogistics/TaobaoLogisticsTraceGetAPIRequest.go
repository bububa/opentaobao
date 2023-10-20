package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsTraceGetAPIRequest 用户根据交易号查询物流流转信息 API请求
// taobao.logistics.trace.get
//
// 用户根据交易号查询物流流转信息
type TaobaoLogisticsTraceGetAPIRequest struct {
	model.Params
	// 交易号
	_tid int64
}

// NewTaobaoLogisticsTraceGetRequest 初始化TaobaoLogisticsTraceGetAPIRequest对象
func NewTaobaoLogisticsTraceGetRequest() *TaobaoLogisticsTraceGetAPIRequest {
	return &TaobaoLogisticsTraceGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsTraceGetAPIRequest) Reset() {
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsTraceGetAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.trace.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsTraceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsTraceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 交易号
func (r *TaobaoLogisticsTraceGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoLogisticsTraceGetAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoLogisticsTraceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsTraceGetRequest()
	},
}

// GetTaobaoLogisticsTraceGetRequest 从 sync.Pool 获取 TaobaoLogisticsTraceGetAPIRequest
func GetTaobaoLogisticsTraceGetAPIRequest() *TaobaoLogisticsTraceGetAPIRequest {
	return poolTaobaoLogisticsTraceGetAPIRequest.Get().(*TaobaoLogisticsTraceGetAPIRequest)
}

// ReleaseTaobaoLogisticsTraceGetAPIRequest 将 TaobaoLogisticsTraceGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsTraceGetAPIRequest(v *TaobaoLogisticsTraceGetAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsTraceGetAPIRequest.Put(v)
}
