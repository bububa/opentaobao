package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkScmLrpOrderPredictAPIRequest 单量预测 API请求
// alibaba.wdk.scm.lrp.order.predict
//
// 提供基于门店和配送站的订单量预测，可用于提前安排人力资源
type AlibabaWdkScmLrpOrderPredictAPIRequest struct {
	model.Params
	// 单量预测查询参数
	_paramOrderPredictQuery *OrderPredictQuery
}

// NewAlibabaWdkScmLrpOrderPredictRequest 初始化AlibabaWdkScmLrpOrderPredictAPIRequest对象
func NewAlibabaWdkScmLrpOrderPredictRequest() *AlibabaWdkScmLrpOrderPredictAPIRequest {
	return &AlibabaWdkScmLrpOrderPredictAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkScmLrpOrderPredictAPIRequest) Reset() {
	r._paramOrderPredictQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkScmLrpOrderPredictAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.scm.lrp.order.predict"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkScmLrpOrderPredictAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkScmLrpOrderPredictAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderPredictQuery is ParamOrderPredictQuery Setter
// 单量预测查询参数
func (r *AlibabaWdkScmLrpOrderPredictAPIRequest) SetParamOrderPredictQuery(_paramOrderPredictQuery *OrderPredictQuery) error {
	r._paramOrderPredictQuery = _paramOrderPredictQuery
	r.Set("param_order_predict_query", _paramOrderPredictQuery)
	return nil
}

// GetParamOrderPredictQuery ParamOrderPredictQuery Getter
func (r AlibabaWdkScmLrpOrderPredictAPIRequest) GetParamOrderPredictQuery() *OrderPredictQuery {
	return r._paramOrderPredictQuery
}

var poolAlibabaWdkScmLrpOrderPredictAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkScmLrpOrderPredictRequest()
	},
}

// GetAlibabaWdkScmLrpOrderPredictRequest 从 sync.Pool 获取 AlibabaWdkScmLrpOrderPredictAPIRequest
func GetAlibabaWdkScmLrpOrderPredictAPIRequest() *AlibabaWdkScmLrpOrderPredictAPIRequest {
	return poolAlibabaWdkScmLrpOrderPredictAPIRequest.Get().(*AlibabaWdkScmLrpOrderPredictAPIRequest)
}

// ReleaseAlibabaWdkScmLrpOrderPredictAPIRequest 将 AlibabaWdkScmLrpOrderPredictAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkScmLrpOrderPredictAPIRequest(v *AlibabaWdkScmLrpOrderPredictAPIRequest) {
	v.Reset()
	poolAlibabaWdkScmLrpOrderPredictAPIRequest.Put(v)
}
