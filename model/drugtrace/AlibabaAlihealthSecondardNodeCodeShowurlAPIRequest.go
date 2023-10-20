package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthsecondardnodecodeshowurlAPIRequest 查询码信息url API请求
// alibaba.alihealth.secondard.node.code.showurl
//
// 二级节点查询码信息url
type AlibabaalihealthsecondardnodecodeshowurlAPIRequest struct {
	model.Params
	// 二级节点用户名
	_userName string
	// 追溯码
	_code string
}

// NewAlibabaalihealthsecondardnodecodeshowurlRequest 初始化AlibabaalihealthsecondardnodecodeshowurlAPIRequest对象
func NewAlibabaalihealthsecondardnodecodeshowurlRequest() *AlibabaalihealthsecondardnodecodeshowurlAPIRequest {
	return &AlibabaalihealthsecondardnodecodeshowurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthsecondardnodecodeshowurlAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.secondard.node.code.showurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthsecondardnodecodeshowurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthsecondardnodecodeshowurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserName is UserName Setter
// 二级节点用户名
func (r *AlibabaalihealthsecondardnodecodeshowurlAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// GetUserName UserName Getter
func (r AlibabaalihealthsecondardnodecodeshowurlAPIRequest) GetUserName() string {
	return r._userName
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthsecondardnodecodeshowurlAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthsecondardnodecodeshowurlAPIRequest) GetCode() string {
	return r._code
}
