package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthUserBaseinfoGetAPIRequest
获取用户基础信息 API请求
alibaba.alihealth.user.baseinfo.get

获取用户基础信息 */
type AlibabaAlihealthUserBaseinfoGetAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
	// 三方服务商
	_appName string
}

// NewAlibabaAlihealthUserBaseinfoGetRequest 初始化AlibabaAlihealthUserBaseinfoGetAPIRequest对象
func NewAlibabaAlihealthUserBaseinfoGetRequest() *AlibabaAlihealthUserBaseinfoGetAPIRequest {
	return &AlibabaAlihealthUserBaseinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthUserBaseinfoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.user.baseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthUserBaseinfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserId Setter
// 用户id
func (r *AlibabaAlihealthUserBaseinfoGetAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaAlihealthUserBaseinfoGetAPIRequest) GetUserId() int64 {
	return r._userId
}

// Set is AppName Setter
// 三方服务商
func (r *AlibabaAlihealthUserBaseinfoGetAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// Get AppName Getter
func (r AlibabaAlihealthUserBaseinfoGetAPIRequest) GetAppName() string {
	return r._appName
}
