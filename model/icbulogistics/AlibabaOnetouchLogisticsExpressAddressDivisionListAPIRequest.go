package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest
四级地址库-区域 API请求
alibaba.onetouch.logistics.express.address.division.list

四级地址库-区 */
type AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// NewAlibabaOnetouchLogisticsExpressAddressDivisionListRequest 初始化AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressDivisionListRequest() *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest {
	return &AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.address.division.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// Get ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) GetParamQuery() *AddressQueryDto {
	return r._paramQuery
}
