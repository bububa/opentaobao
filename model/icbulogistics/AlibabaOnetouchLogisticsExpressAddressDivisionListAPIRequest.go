package icbulogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest 四级地址库-区域 API请求
// alibaba.onetouch.logistics.express.address.division.list
//
// 四级地址库-区
type AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// NewAlibabaOnetouchLogisticsExpressAddressDivisionListRequest 初始化AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressDivisionListRequest() *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest {
	return &AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) Reset() {
	r._paramQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.address.division.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) GetParamQuery() *AddressQueryDto {
	return r._paramQuery
}

var poolAlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOnetouchLogisticsExpressAddressDivisionListRequest()
	},
}

// GetAlibabaOnetouchLogisticsExpressAddressDivisionListRequest 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest
func GetAlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest() *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest {
	return poolAlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest.Get().(*AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest)
}

// ReleaseAlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest 将 AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest 放入 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest(v *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest.Put(v)
}
