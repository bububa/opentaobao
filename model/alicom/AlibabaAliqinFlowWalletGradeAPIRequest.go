package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletGradeAPIRequest 获取流量档位 API请求
// alibaba.aliqin.flow.wallet.grade
//
// 获取直充流量档位
type AlibabaAliqinFlowWalletGradeAPIRequest struct {
	model.Params
	// 手机号码
	_phoneNum string
}

// NewAlibabaAliqinFlowWalletGradeRequest 初始化AlibabaAliqinFlowWalletGradeAPIRequest对象
func NewAlibabaAliqinFlowWalletGradeRequest() *AlibabaAliqinFlowWalletGradeAPIRequest {
	return &AlibabaAliqinFlowWalletGradeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFlowWalletGradeAPIRequest) Reset() {
	r._phoneNum = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletGradeAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.grade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletGradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFlowWalletGradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhoneNum is PhoneNum Setter
// 手机号码
func (r *AlibabaAliqinFlowWalletGradeAPIRequest) SetPhoneNum(_phoneNum string) error {
	r._phoneNum = _phoneNum
	r.Set("phone_num", _phoneNum)
	return nil
}

// GetPhoneNum PhoneNum Getter
func (r AlibabaAliqinFlowWalletGradeAPIRequest) GetPhoneNum() string {
	return r._phoneNum
}

var poolAlibabaAliqinFlowWalletGradeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFlowWalletGradeRequest()
	},
}

// GetAlibabaAliqinFlowWalletGradeRequest 从 sync.Pool 获取 AlibabaAliqinFlowWalletGradeAPIRequest
func GetAlibabaAliqinFlowWalletGradeAPIRequest() *AlibabaAliqinFlowWalletGradeAPIRequest {
	return poolAlibabaAliqinFlowWalletGradeAPIRequest.Get().(*AlibabaAliqinFlowWalletGradeAPIRequest)
}

// ReleaseAlibabaAliqinFlowWalletGradeAPIRequest 将 AlibabaAliqinFlowWalletGradeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFlowWalletGradeAPIRequest(v *AlibabaAliqinFlowWalletGradeAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFlowWalletGradeAPIRequest.Put(v)
}
