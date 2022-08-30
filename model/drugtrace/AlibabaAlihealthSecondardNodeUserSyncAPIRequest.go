package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthSecondardNodeUserSyncAPIRequest 二级节点用户数据同步 API请求
// alibaba.alihealth.secondard.node.user.sync
//
// 二级节点用户数据同步
type AlibabaAlihealthSecondardNodeUserSyncAPIRequest struct {
	model.Params
	// 用户名
	_userName string
	// 手机号
	_mobile string
	// 邮箱
	_email string
	// 注册时间，yyyy-MM-dd
	_regDate string
}

// NewAlibabaAlihealthSecondardNodeUserSyncRequest 初始化AlibabaAlihealthSecondardNodeUserSyncAPIRequest对象
func NewAlibabaAlihealthSecondardNodeUserSyncRequest() *AlibabaAlihealthSecondardNodeUserSyncAPIRequest {
	return &AlibabaAlihealthSecondardNodeUserSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthSecondardNodeUserSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.secondard.node.user.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthSecondardNodeUserSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserName is UserName Setter
// 用户名
func (r *AlibabaAlihealthSecondardNodeUserSyncAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// GetUserName UserName Getter
func (r AlibabaAlihealthSecondardNodeUserSyncAPIRequest) GetUserName() string {
	return r._userName
}

// SetMobile is Mobile Setter
// 手机号
func (r *AlibabaAlihealthSecondardNodeUserSyncAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaAlihealthSecondardNodeUserSyncAPIRequest) GetMobile() string {
	return r._mobile
}

// SetEmail is Email Setter
// 邮箱
func (r *AlibabaAlihealthSecondardNodeUserSyncAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r AlibabaAlihealthSecondardNodeUserSyncAPIRequest) GetEmail() string {
	return r._email
}

// SetRegDate is RegDate Setter
// 注册时间，yyyy-MM-dd
func (r *AlibabaAlihealthSecondardNodeUserSyncAPIRequest) SetRegDate(_regDate string) error {
	r._regDate = _regDate
	r.Set("reg_date", _regDate)
	return nil
}

// GetRegDate RegDate Getter
func (r AlibabaAlihealthSecondardNodeUserSyncAPIRequest) GetRegDate() string {
	return r._regDate
}
