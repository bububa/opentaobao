package icbulogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest 四级地址库-市 API请求
// alibaba.onetouch.logistics.express.address.city.list
//
// 四级地址库-市
type AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// NewAlibabaOnetouchLogisticsExpressAddressCityListRequest 初始化AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressCityListRequest() *AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest {
	return &AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) Reset() {
	r._paramQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.address.city.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) GetParamQuery() *AddressQueryDto {
	return r._paramQuery
}

var poolAlibabaOnetouchLogisticsExpressAddressCityListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOnetouchLogisticsExpressAddressCityListRequest()
	},
}

// GetAlibabaOnetouchLogisticsExpressAddressCityListRequest 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest
func GetAlibabaOnetouchLogisticsExpressAddressCityListAPIRequest() *AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest {
	return poolAlibabaOnetouchLogisticsExpressAddressCityListAPIRequest.Get().(*AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest)
}

// ReleaseAlibabaOnetouchLogisticsExpressAddressCityListAPIRequest 将 AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest 放入 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressAddressCityListAPIRequest(v *AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressAddressCityListAPIRequest.Put(v)
}
