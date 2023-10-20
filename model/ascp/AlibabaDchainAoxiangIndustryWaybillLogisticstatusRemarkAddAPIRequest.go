package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest 客服增加备注 API请求
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.remark.add
//
// 客服增加备注
type AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest struct {
	model.Params
	// 物流节点回传入参
	_tmsOrderRemarkRequest *TmsOrderRemarkRequest
}

// NewAlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddRequest 初始化AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest对象
func NewAlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddRequest() *AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest {
	return &AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.logisticstatus.remark.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsOrderRemarkRequest is TmsOrderRemarkRequest Setter
// 物流节点回传入参
func (r *AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest) SetTmsOrderRemarkRequest(_tmsOrderRemarkRequest *TmsOrderRemarkRequest) error {
	r._tmsOrderRemarkRequest = _tmsOrderRemarkRequest
	r.Set("tms_order_remark_request", _tmsOrderRemarkRequest)
	return nil
}

// GetTmsOrderRemarkRequest TmsOrderRemarkRequest Getter
func (r AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest) GetTmsOrderRemarkRequest() *TmsOrderRemarkRequest {
	return r._tmsOrderRemarkRequest
}
