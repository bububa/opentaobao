package user

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServiceSettleadjustmentModifyAPIRequest) Reset() {
	r._paramSettleAdjustmentRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentModifyAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServiceSettleadjustmentModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSettleAdjustmentRequest is ParamSettleAdjustmentRequest Setter
// 结算调整单父节点
func (r *TmallServiceSettleadjustmentModifyAPIRequest) SetParamSettleAdjustmentRequest(_paramSettleAdjustmentRequest *SettleAdjustmentRequest) error {
	r._paramSettleAdjustmentRequest = _paramSettleAdjustmentRequest
	r.Set("param_settle_adjustment_request", _paramSettleAdjustmentRequest)
	return nil
}

// GetParamSettleAdjustmentRequest ParamSettleAdjustmentRequest Getter
func (r TmallServiceSettleadjustmentModifyAPIRequest) GetParamSettleAdjustmentRequest() *SettleAdjustmentRequest {
	return r._paramSettleAdjustmentRequest
}

var poolTmallServiceSettleadjustmentModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServiceSettleadjustmentModifyRequest()
	},
}

// GetTmallServiceSettleadjustmentModifyRequest 从 sync.Pool 获取 TmallServiceSettleadjustmentModifyAPIRequest
func GetTmallServiceSettleadjustmentModifyAPIRequest() *TmallServiceSettleadjustmentModifyAPIRequest {
	return poolTmallServiceSettleadjustmentModifyAPIRequest.Get().(*TmallServiceSettleadjustmentModifyAPIRequest)
}

// ReleaseTmallServiceSettleadjustmentModifyAPIRequest 将 TmallServiceSettleadjustmentModifyAPIRequest 放入 sync.Pool
func ReleaseTmallServiceSettleadjustmentModifyAPIRequest(v *TmallServiceSettleadjustmentModifyAPIRequest) {
	v.Reset()
	poolTmallServiceSettleadjustmentModifyAPIRequest.Put(v)
}
