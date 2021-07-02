package icbulogistics

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.address.province.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
