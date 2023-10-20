package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangindustrywaybillcreateAPIRequest 服务商开运单 API请求
// alibaba.dchain.aoxiang.industry.waybill.create
//
// 服务商开运单
type AlibabadchainaoxiangindustrywaybillcreateAPIRequest struct {
	model.Params
	// 服务商开单
	_tmsOrderCreateRequest *TmsOrderCreateRequest
}

// NewAlibabadchainaoxiangindustrywaybillcreateRequest 初始化AlibabadchainaoxiangindustrywaybillcreateAPIRequest对象
func NewAlibabadchainaoxiangindustrywaybillcreateRequest() *AlibabadchainaoxiangindustrywaybillcreateAPIRequest {
	return &AlibabadchainaoxiangindustrywaybillcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangindustrywaybillcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangindustrywaybillcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangindustrywaybillcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsOrderCreateRequest is TmsOrderCreateRequest Setter
// 服务商开单
func (r *AlibabadchainaoxiangindustrywaybillcreateAPIRequest) SetTmsOrderCreateRequest(_tmsOrderCreateRequest *TmsOrderCreateRequest) error {
	r._tmsOrderCreateRequest = _tmsOrderCreateRequest
	r.Set("tms_order_create_request", _tmsOrderCreateRequest)
	return nil
}

// GetTmsOrderCreateRequest TmsOrderCreateRequest Getter
func (r AlibabadchainaoxiangindustrywaybillcreateAPIRequest) GetTmsOrderCreateRequest() *TmsOrderCreateRequest {
	return r._tmsOrderCreateRequest
}
