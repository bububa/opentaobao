package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentRequestAPIRequest 创建结算调整单 API请求
// tmall.service.settleadjustment.request
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明工单ID，调整费用值，调整原因进行新建结算调整单。
type TmallServiceSettleadjustmentRequestAPIRequest struct {
	model.Params
	// 父节点
	_paramSettleAdjustmentRequest *SettleAdjustmentRequest
}

// NewTmallServiceSettleadjustmentRequestRequest 初始化TmallServiceSettleadjustmentRequestAPIRequest对象
func NewTmallServiceSettleadjustmentRequestRequest() *TmallServiceSettleadjustmentRequestAPIRequest {
	return &TmallServiceSettleadjustmentRequestAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServiceSettleadjustmentRequestAPIRequest) Reset() {
	r._paramSettleAdjustmentRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentRequestAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.request"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentRequestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServiceSettleadjustmentRequestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSettleAdjustmentRequest is ParamSettleAdjustmentRequest Setter
// 父节点
func (r *TmallServiceSettleadjustmentRequestAPIRequest) SetParamSettleAdjustmentRequest(_paramSettleAdjustmentRequest *SettleAdjustmentRequest) error {
	r._paramSettleAdjustmentRequest = _paramSettleAdjustmentRequest
	r.Set("param_settle_adjustment_request", _paramSettleAdjustmentRequest)
	return nil
}

// GetParamSettleAdjustmentRequest ParamSettleAdjustmentRequest Getter
func (r TmallServiceSettleadjustmentRequestAPIRequest) GetParamSettleAdjustmentRequest() *SettleAdjustmentRequest {
	return r._paramSettleAdjustmentRequest
}

var poolTmallServiceSettleadjustmentRequestAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServiceSettleadjustmentRequestRequest()
	},
}

// GetTmallServiceSettleadjustmentRequestRequest 从 sync.Pool 获取 TmallServiceSettleadjustmentRequestAPIRequest
func GetTmallServiceSettleadjustmentRequestAPIRequest() *TmallServiceSettleadjustmentRequestAPIRequest {
	return poolTmallServiceSettleadjustmentRequestAPIRequest.Get().(*TmallServiceSettleadjustmentRequestAPIRequest)
}

// ReleaseTmallServiceSettleadjustmentRequestAPIRequest 将 TmallServiceSettleadjustmentRequestAPIRequest 放入 sync.Pool
func ReleaseTmallServiceSettleadjustmentRequestAPIRequest(v *TmallServiceSettleadjustmentRequestAPIRequest) {
	v.Reset()
	poolTmallServiceSettleadjustmentRequestAPIRequest.Put(v)
}
