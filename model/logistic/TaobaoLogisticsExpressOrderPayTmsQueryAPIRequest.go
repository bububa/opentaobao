package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest 上门取退运费支付状态查询接口 API请求
// taobao.logistics.express.order.pay.tms.query
//
// 上门取退运费支付状态查询接口
type TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest struct {
	model.Params
	// 查询入参
	_tms2MscPayQueryRequest *Tms2MscPayQueryRequest
}

// NewTaobaoLogisticsExpressOrderPayTmsQueryRequest 初始化TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest对象
func NewTaobaoLogisticsExpressOrderPayTmsQueryRequest() *TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest {
	return &TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest) Reset() {
	r._tms2MscPayQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.order.pay.tms.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTms2MscPayQueryRequest is Tms2MscPayQueryRequest Setter
// 查询入参
func (r *TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest) SetTms2MscPayQueryRequest(_tms2MscPayQueryRequest *Tms2MscPayQueryRequest) error {
	r._tms2MscPayQueryRequest = _tms2MscPayQueryRequest
	r.Set("tms2_msc_pay_query_request", _tms2MscPayQueryRequest)
	return nil
}

// GetTms2MscPayQueryRequest Tms2MscPayQueryRequest Getter
func (r TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest) GetTms2MscPayQueryRequest() *Tms2MscPayQueryRequest {
	return r._tms2MscPayQueryRequest
}

var poolTaobaoLogisticsExpressOrderPayTmsQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressOrderPayTmsQueryRequest()
	},
}

// GetTaobaoLogisticsExpressOrderPayTmsQueryRequest 从 sync.Pool 获取 TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest
func GetTaobaoLogisticsExpressOrderPayTmsQueryAPIRequest() *TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest {
	return poolTaobaoLogisticsExpressOrderPayTmsQueryAPIRequest.Get().(*TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest)
}

// ReleaseTaobaoLogisticsExpressOrderPayTmsQueryAPIRequest 将 TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressOrderPayTmsQueryAPIRequest(v *TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressOrderPayTmsQueryAPIRequest.Put(v)
}
