package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest 查询码信息url API请求
// alibaba.alihealth.secondard.node.code.showurl
//
// 二级节点查询码信息url
type AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest struct {
	model.Params
	// 二级节点用户名
	_userName string
	// 追溯码
	_code string
}

// NewAlibabaAlihealthSecondardNodeCodeShowurlRequest 初始化AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest对象
func NewAlibabaAlihealthSecondardNodeCodeShowurlRequest() *AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest {
	return &AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) Reset() {
	r._userName = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.secondard.node.code.showurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserName is UserName Setter
// 二级节点用户名
func (r *AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// GetUserName UserName Getter
func (r AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) GetUserName() string {
	return r._userName
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaAlihealthSecondardNodeCodeShowurlAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthSecondardNodeCodeShowurlRequest()
	},
}

// GetAlibabaAlihealthSecondardNodeCodeShowurlRequest 从 sync.Pool 获取 AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest
func GetAlibabaAlihealthSecondardNodeCodeShowurlAPIRequest() *AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest {
	return poolAlibabaAlihealthSecondardNodeCodeShowurlAPIRequest.Get().(*AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest)
}

// ReleaseAlibabaAlihealthSecondardNodeCodeShowurlAPIRequest 将 AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthSecondardNodeCodeShowurlAPIRequest(v *AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthSecondardNodeCodeShowurlAPIRequest.Put(v)
}
