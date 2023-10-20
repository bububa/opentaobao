package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrechargeunchargecheckgetAPIRequest 储值账户退充值校验 API请求
// alibaba.alsc.crm.recharge.unchargecheck.get
//
// 储值账户退充值校验接口
type AlibabaalsccrmrechargeunchargecheckgetAPIRequest struct {
	model.Params
	// 入参
	_paramUnchargeCheckOpenReq *UnchargeCheckOpenReq
}

// NewAlibabaalsccrmrechargeunchargecheckgetRequest 初始化AlibabaalsccrmrechargeunchargecheckgetAPIRequest对象
func NewAlibabaalsccrmrechargeunchargecheckgetRequest() *AlibabaalsccrmrechargeunchargecheckgetAPIRequest {
	return &AlibabaalsccrmrechargeunchargecheckgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrechargeunchargecheckgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.unchargecheck.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrechargeunchargecheckgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrechargeunchargecheckgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUnchargeCheckOpenReq is ParamUnchargeCheckOpenReq Setter
// 入参
func (r *AlibabaalsccrmrechargeunchargecheckgetAPIRequest) SetParamUnchargeCheckOpenReq(_paramUnchargeCheckOpenReq *UnchargeCheckOpenReq) error {
	r._paramUnchargeCheckOpenReq = _paramUnchargeCheckOpenReq
	r.Set("param_uncharge_check_open_req", _paramUnchargeCheckOpenReq)
	return nil
}

// GetParamUnchargeCheckOpenReq ParamUnchargeCheckOpenReq Getter
func (r AlibabaalsccrmrechargeunchargecheckgetAPIRequest) GetParamUnchargeCheckOpenReq() *UnchargeCheckOpenReq {
	return r._paramUnchargeCheckOpenReq
}
