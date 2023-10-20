package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrechargechargeupdateAPIRequest 储值充值 API请求
// alibaba.alsc.crm.recharge.charge.update
//
// 顾客储值账户充值
type AlibabaalsccrmrechargechargeupdateAPIRequest struct {
	model.Params
	// 入参
	_paramRechargeOpenReq *RechargeOpenReq
}

// NewAlibabaalsccrmrechargechargeupdateRequest 初始化AlibabaalsccrmrechargechargeupdateAPIRequest对象
func NewAlibabaalsccrmrechargechargeupdateRequest() *AlibabaalsccrmrechargechargeupdateAPIRequest {
	return &AlibabaalsccrmrechargechargeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrechargechargeupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.charge.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrechargechargeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrechargechargeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRechargeOpenReq is ParamRechargeOpenReq Setter
// 入参
func (r *AlibabaalsccrmrechargechargeupdateAPIRequest) SetParamRechargeOpenReq(_paramRechargeOpenReq *RechargeOpenReq) error {
	r._paramRechargeOpenReq = _paramRechargeOpenReq
	r.Set("param_recharge_open_req", _paramRechargeOpenReq)
	return nil
}

// GetParamRechargeOpenReq ParamRechargeOpenReq Getter
func (r AlibabaalsccrmrechargechargeupdateAPIRequest) GetParamRechargeOpenReq() *RechargeOpenReq {
	return r._paramRechargeOpenReq
}
