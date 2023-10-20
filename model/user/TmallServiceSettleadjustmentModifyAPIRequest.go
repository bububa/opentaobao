package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicesettleadjustmentmodifyAPIRequest 修改结算调整单 API请求
// tmall.service.settleadjustment.modify
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
type TmallservicesettleadjustmentmodifyAPIRequest struct {
	model.Params
	// 结算调整单父节点
	_paramSettleAdjustmentRequest *SettleAdjustmentRequest
}

// NewTmallservicesettleadjustmentmodifyRequest 初始化TmallservicesettleadjustmentmodifyAPIRequest对象
func NewTmallservicesettleadjustmentmodifyRequest() *TmallservicesettleadjustmentmodifyAPIRequest {
	return &TmallservicesettleadjustmentmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicesettleadjustmentmodifyAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicesettleadjustmentmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicesettleadjustmentmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSettleAdjustmentRequest is ParamSettleAdjustmentRequest Setter
// 结算调整单父节点
func (r *TmallservicesettleadjustmentmodifyAPIRequest) SetParamSettleAdjustmentRequest(_paramSettleAdjustmentRequest *SettleAdjustmentRequest) error {
	r._paramSettleAdjustmentRequest = _paramSettleAdjustmentRequest
	r.Set("param_settle_adjustment_request", _paramSettleAdjustmentRequest)
	return nil
}

// GetParamSettleAdjustmentRequest ParamSettleAdjustmentRequest Getter
func (r TmallservicesettleadjustmentmodifyAPIRequest) GetParamSettleAdjustmentRequest() *SettleAdjustmentRequest {
	return r._paramSettleAdjustmentRequest
}
