package drugtrace

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.secondard.node.code.showurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
