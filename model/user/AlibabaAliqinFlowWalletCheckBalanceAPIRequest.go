package user

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFlowWalletCheckBalanceAPIRequest) Reset() {
	r._gradeId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletCheckBalanceAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.check.balance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletCheckBalanceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFlowWalletCheckBalanceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGradeId is GradeId Setter
// 检查金额档位id
func (r *AlibabaAliqinFlowWalletCheckBalanceAPIRequest) SetGradeId(_gradeId string) error {
	r._gradeId = _gradeId
	r.Set("grade_id", _gradeId)
	return nil
}

// GetGradeId GradeId Getter
func (r AlibabaAliqinFlowWalletCheckBalanceAPIRequest) GetGradeId() string {
	return r._gradeId
}

var poolAlibabaAliqinFlowWalletCheckBalanceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFlowWalletCheckBalanceRequest()
	},
}

// GetAlibabaAliqinFlowWalletCheckBalanceRequest 从 sync.Pool 获取 AlibabaAliqinFlowWalletCheckBalanceAPIRequest
func GetAlibabaAliqinFlowWalletCheckBalanceAPIRequest() *AlibabaAliqinFlowWalletCheckBalanceAPIRequest {
	return poolAlibabaAliqinFlowWalletCheckBalanceAPIRequest.Get().(*AlibabaAliqinFlowWalletCheckBalanceAPIRequest)
}

// ReleaseAlibabaAliqinFlowWalletCheckBalanceAPIRequest 将 AlibabaAliqinFlowWalletCheckBalanceAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFlowWalletCheckBalanceAPIRequest(v *AlibabaAliqinFlowWalletCheckBalanceAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFlowWalletCheckBalanceAPIRequest.Put(v)
}
