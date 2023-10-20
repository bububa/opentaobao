package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest 物流节点回传 API请求
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.change
//
// 物流节点回传
type AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest struct {
	model.Params
	// 物流节点回传入参
	_tmsOrderConfirmRequest *TmsOrderConfirmRequest
}

// NewAlibabadchainaoxiangindustrywaybilllogisticstatuschangeRequest 初始化AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest对象
func NewAlibabadchainaoxiangindustrywaybilllogisticstatuschangeRequest() *AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest {
	return &AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.logisticstatus.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsOrderConfirmRequest is TmsOrderConfirmRequest Setter
// 物流节点回传入参
func (r *AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest) SetTmsOrderConfirmRequest(_tmsOrderConfirmRequest *TmsOrderConfirmRequest) error {
	r._tmsOrderConfirmRequest = _tmsOrderConfirmRequest
	r.Set("tms_order_confirm_request", _tmsOrderConfirmRequest)
	return nil
}

// GetTmsOrderConfirmRequest TmsOrderConfirmRequest Getter
func (r AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest) GetTmsOrderConfirmRequest() *TmsOrderConfirmRequest {
	return r._tmsOrderConfirmRequest
}
