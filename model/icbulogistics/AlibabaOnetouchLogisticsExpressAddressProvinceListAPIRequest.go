package icbulogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest 四级地址库-省 API请求
// alibaba.onetouch.logistics.express.address.province.list
//
// 四级地址库-省
type AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// NewAlibabaOnetouchLogisticsExpressAddressProvinceListRequest 初始化AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressProvinceListRequest() *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest {
	return &AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) Reset() {
	r._paramQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.address.province.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) GetParamQuery() *AddressQueryDto {
	return r._paramQuery
}

var poolAlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOnetouchLogisticsExpressAddressProvinceListRequest()
	},
}

// GetAlibabaOnetouchLogisticsExpressAddressProvinceListRequest 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest
func GetAlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest() *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest {
	return poolAlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest.Get().(*AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest)
}

// ReleaseAlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest 将 AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest 放入 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest(v *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest.Put(v)
}
