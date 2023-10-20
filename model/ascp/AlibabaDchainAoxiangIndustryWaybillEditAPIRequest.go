package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangindustrywaybilleditAPIRequest 服务商编辑运单 API请求
// alibaba.dchain.aoxiang.industry.waybill.edit
//
// 服务商编辑运单
type AlibabadchainaoxiangindustrywaybilleditAPIRequest struct {
	model.Params
	// 服务商编辑
	_tmsOrderUpdateRequest *TmsOrderUpdateRequest
}

// NewAlibabadchainaoxiangindustrywaybilleditRequest 初始化AlibabadchainaoxiangindustrywaybilleditAPIRequest对象
func NewAlibabadchainaoxiangindustrywaybilleditRequest() *AlibabadchainaoxiangindustrywaybilleditAPIRequest {
	return &AlibabadchainaoxiangindustrywaybilleditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangindustrywaybilleditAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangindustrywaybilleditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangindustrywaybilleditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsOrderUpdateRequest is TmsOrderUpdateRequest Setter
// 服务商编辑
func (r *AlibabadchainaoxiangindustrywaybilleditAPIRequest) SetTmsOrderUpdateRequest(_tmsOrderUpdateRequest *TmsOrderUpdateRequest) error {
	r._tmsOrderUpdateRequest = _tmsOrderUpdateRequest
	r.Set("tms_order_update_request", _tmsOrderUpdateRequest)
	return nil
}

// GetTmsOrderUpdateRequest TmsOrderUpdateRequest Getter
func (r AlibabadchainaoxiangindustrywaybilleditAPIRequest) GetTmsOrderUpdateRequest() *TmsOrderUpdateRequest {
	return r._tmsOrderUpdateRequest
}
