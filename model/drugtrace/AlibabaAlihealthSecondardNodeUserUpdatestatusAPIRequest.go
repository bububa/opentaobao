package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthSecondardNodeUserUpdatestatusAPIRequest 二级节点用户注销 API请求
// alibaba.alihealth.secondard.node.user.updatestatus
//
// 注销二级节点用户
type AlibabaAlihealthSecondardNodeUserUpdatestatusAPIRequest struct {
	model.Params
	// 用户名
	_userName string
}

// NewAlibabaAlihealthSecondardNodeUserUpdatestatusRequest 初始化AlibabaAlihealthSecondardNodeUserUpdatestatusAPIRequest对象
func NewAlibabaAlihealthSecondardNodeUserUpdatestatusRequest() *AlibabaAlihealthSecondardNodeUserUpdatestatusAPIRequest {
	return &AlibabaAlihealthSecondardNodeUserUpdatestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthSecondardNodeUserUpdatestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.secondard.node.user.updatestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthSecondardNodeUserUpdatestatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserName is UserName Setter
// 用户名
func (r *AlibabaAlihealthSecondardNodeUserUpdatestatusAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// GetUserName UserName Getter
func (r AlibabaAlihealthSecondardNodeUserUpdatestatusAPIRequest) GetUserName() string {
	return r._userName
}
