package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicesettleadjustmentrequestAPIRequest 创建结算调整单 API请求
// tmall.service.settleadjustment.request
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明工单ID，调整费用值，调整原因进行新建结算调整单。
type TmallservicesettleadjustmentrequestAPIRequest struct {
	model.Params
	// 父节点
	_paramSettleAdjustmentRequest *SettleAdjustmentRequest
}

// NewTmallservicesettleadjustmentrequestRequest 初始化TmallservicesettleadjustmentrequestAPIRequest对象
func NewTmallservicesettleadjustmentrequestRequest() *TmallservicesettleadjustmentrequestAPIRequest {
	return &TmallservicesettleadjustmentrequestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicesettleadjustmentrequestAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.request"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicesettleadjustmentrequestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicesettleadjustmentrequestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSettleAdjustmentRequest is ParamSettleAdjustmentRequest Setter
// 父节点
func (r *TmallservicesettleadjustmentrequestAPIRequest) SetParamSettleAdjustmentRequest(_paramSettleAdjustmentRequest *SettleAdjustmentRequest) error {
	r._paramSettleAdjustmentRequest = _paramSettleAdjustmentRequest
	r.Set("param_settle_adjustment_request", _paramSettleAdjustmentRequest)
	return nil
}

// GetParamSettleAdjustmentRequest ParamSettleAdjustmentRequest Getter
func (r TmallservicesettleadjustmentrequestAPIRequest) GetParamSettleAdjustmentRequest() *SettleAdjustmentRequest {
	return r._paramSettleAdjustmentRequest
}
