package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletCheckBalanceAPIRequest 商家预存余额检查 API请求
// alibaba.aliqin.flow.wallet.check.balance
//
// 检查商家CRM预存余额是否足够进行活动
type AlibabaAliqinFlowWalletCheckBalanceAPIRequest struct {
	model.Params
	// 检查金额档位id
	_gradeId string
}

// NewAlibabaAliqinFlowWalletCheckBalanceRequest 初始化AlibabaAliqinFlowWalletCheckBalanceAPIRequest对象
func NewAlibabaAliqinFlowWalletCheckBalanceRequest() *AlibabaAliqinFlowWalletCheckBalanceAPIRequest {
	return &AlibabaAliqinFlowWalletCheckBalanceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletCheckBalanceAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.check.balance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletCheckBalanceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GradeId Setter
// 检查金额档位id
func (r *AlibabaAliqinFlowWalletCheckBalanceAPIRequest) SetGradeId(_gradeId string) error {
	r._gradeId = _gradeId
	r.Set("grade_id", _gradeId)
	return nil
}

// Get GradeId Getter
func (r AlibabaAliqinFlowWalletCheckBalanceAPIRequest) GetGradeId() string {
	return r._gradeId
}
