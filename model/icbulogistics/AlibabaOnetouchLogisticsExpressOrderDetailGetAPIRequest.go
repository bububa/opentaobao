package icbulogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest 订单详细信息(面单及仓库信息) API请求
// alibaba.onetouch.logistics.express.order.detail.get
//
// 订单详细信息(面单及仓库信息)
type AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *LogisticsOrderQueryDto
}

// NewAlibabaOnetouchLogisticsExpressOrderDetailGetRequest 初始化AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressOrderDetailGetRequest() *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest {
	return &AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) Reset() {
	r._paramQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.order.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) SetParamQuery(_paramQuery *LogisticsOrderQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) GetParamQuery() *LogisticsOrderQueryDto {
	return r._paramQuery
}

var poolAlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOnetouchLogisticsExpressOrderDetailGetRequest()
	},
}

// GetAlibabaOnetouchLogisticsExpressOrderDetailGetRequest 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest
func GetAlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest() *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest {
	return poolAlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest.Get().(*AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest)
}

// ReleaseAlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest 将 AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest(v *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest.Put(v)
}
