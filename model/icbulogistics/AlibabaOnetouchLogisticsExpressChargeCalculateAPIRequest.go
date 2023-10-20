package icbulogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest 计算快递运费&下单参数校验 API请求
// alibaba.onetouch.logistics.express.charge.calculate
//
// 计算快递运费、下单参数校验
type AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest struct {
	model.Params
	// 请求参数对象
	_paramnQuery *PlaceOrderDto
}

// NewAlibabaOnetouchLogisticsExpressChargeCalculateRequest 初始化AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressChargeCalculateRequest() *AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest {
	return &AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) Reset() {
	r._paramnQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.charge.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamnQuery is ParamnQuery Setter
// 请求参数对象
func (r *AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) SetParamnQuery(_paramnQuery *PlaceOrderDto) error {
	r._paramnQuery = _paramnQuery
	r.Set("paramn_query", _paramnQuery)
	return nil
}

// GetParamnQuery ParamnQuery Getter
func (r AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) GetParamnQuery() *PlaceOrderDto {
	return r._paramnQuery
}

var poolAlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOnetouchLogisticsExpressChargeCalculateRequest()
	},
}

// GetAlibabaOnetouchLogisticsExpressChargeCalculateRequest 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest
func GetAlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest() *AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest {
	return poolAlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest.Get().(*AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest)
}

// ReleaseAlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest 将 AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest 放入 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest(v *AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest.Put(v)
}
