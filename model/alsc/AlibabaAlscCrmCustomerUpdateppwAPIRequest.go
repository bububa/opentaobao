package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcustomerupdateppwAPIRequest 修改支付密码 API请求
// alibaba.alsc.crm.customer.updateppw
//
// 修改支付密码
type AlibabaalsccrmcustomerupdateppwAPIRequest struct {
	model.Params
	// 修改密码
	_updatePayPasswdReq *UpdatePayPasswdReq
}

// NewAlibabaalsccrmcustomerupdateppwRequest 初始化AlibabaalsccrmcustomerupdateppwAPIRequest对象
func NewAlibabaalsccrmcustomerupdateppwRequest() *AlibabaalsccrmcustomerupdateppwAPIRequest {
	return &AlibabaalsccrmcustomerupdateppwAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcustomerupdateppwAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.updateppw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcustomerupdateppwAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcustomerupdateppwAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdatePayPasswdReq is UpdatePayPasswdReq Setter
// 修改密码
func (r *AlibabaalsccrmcustomerupdateppwAPIRequest) SetUpdatePayPasswdReq(_updatePayPasswdReq *UpdatePayPasswdReq) error {
	r._updatePayPasswdReq = _updatePayPasswdReq
	r.Set("update_pay_passwd_req", _updatePayPasswdReq)
	return nil
}

// GetUpdatePayPasswdReq UpdatePayPasswdReq Getter
func (r AlibabaalsccrmcustomerupdateppwAPIRequest) GetUpdatePayPasswdReq() *UpdatePayPasswdReq {
	return r._updatePayPasswdReq
}
