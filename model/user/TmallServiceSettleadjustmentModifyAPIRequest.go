package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentModifyAPIRequest 修改结算调整单 API请求
// tmall.service.settleadjustment.modify
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
type TmallServiceSettleadjustmentModifyAPIRequest struct {
	model.Params
	// 结算调整单父节点
	_paramSettleAdjustmentRequest *SettleAdjustmentRequest
}

// NewTmallServiceSettleadjustmentModifyRequest 初始化TmallServiceSettleadjustmentModifyAPIRequest对象
func NewTmallServiceSettleadjustmentModifyRequest() *TmallServiceSettleadjustmentModifyAPIRequest {
	return &TmallServiceSettleadjustmentModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentModifyAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamSettleAdjustmentRequest Setter
// 结算调整单父节点
func (r *TmallServiceSettleadjustmentModifyAPIRequest) SetParamSettleAdjustmentRequest(_paramSettleAdjustmentRequest *SettleAdjustmentRequest) error {
	r._paramSettleAdjustmentRequest = _paramSettleAdjustmentRequest
	r.Set("param_settle_adjustment_request", _paramSettleAdjustmentRequest)
	return nil
}

// Get ParamSettleAdjustmentRequest Getter
func (r TmallServiceSettleadjustmentModifyAPIRequest) GetParamSettleAdjustmentRequest() *SettleAdjustmentRequest {
	return r._paramSettleAdjustmentRequest
}
