package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrechargededutprecheckgetAPIRequest 储值核销预先校验 API请求
// alibaba.alsc.crm.recharge.dedutprecheck.get
//
// 储值核销预先校验接口
type AlibabaalsccrmrechargededutprecheckgetAPIRequest struct {
	model.Params
	// 入参
	_paramDeductPreCheckOpenReq *DeductPreCheckOpenReq
}

// NewAlibabaalsccrmrechargededutprecheckgetRequest 初始化AlibabaalsccrmrechargededutprecheckgetAPIRequest对象
func NewAlibabaalsccrmrechargededutprecheckgetRequest() *AlibabaalsccrmrechargededutprecheckgetAPIRequest {
	return &AlibabaalsccrmrechargededutprecheckgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrechargededutprecheckgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.dedutprecheck.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrechargededutprecheckgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrechargededutprecheckgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamDeductPreCheckOpenReq is ParamDeductPreCheckOpenReq Setter
// 入参
func (r *AlibabaalsccrmrechargededutprecheckgetAPIRequest) SetParamDeductPreCheckOpenReq(_paramDeductPreCheckOpenReq *DeductPreCheckOpenReq) error {
	r._paramDeductPreCheckOpenReq = _paramDeductPreCheckOpenReq
	r.Set("param_deduct_pre_check_open_req", _paramDeductPreCheckOpenReq)
	return nil
}

// GetParamDeductPreCheckOpenReq ParamDeductPreCheckOpenReq Getter
func (r AlibabaalsccrmrechargededutprecheckgetAPIRequest) GetParamDeductPreCheckOpenReq() *DeductPreCheckOpenReq {
	return r._paramDeductPreCheckOpenReq
}
