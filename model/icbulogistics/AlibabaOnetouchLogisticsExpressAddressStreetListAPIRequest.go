package icbulogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest 四级地址库-街道 API请求
// alibaba.onetouch.logistics.express.address.street.list
//
// 四级地址库-街道模糊查询
type AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// NewAlibabaOnetouchLogisticsExpressAddressStreetListRequest 初始化AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressStreetListRequest() *AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest {
	return &AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest) Reset() {
	r._paramQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.address.street.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest) GetParamQuery() *AddressQueryDto {
	return r._paramQuery
}

var poolAlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOnetouchLogisticsExpressAddressStreetListRequest()
	},
}

// GetAlibabaOnetouchLogisticsExpressAddressStreetListRequest 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest
func GetAlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest() *AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest {
	return poolAlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest.Get().(*AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest)
}

// ReleaseAlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest 将 AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest 放入 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest(v *AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest.Put(v)
}
