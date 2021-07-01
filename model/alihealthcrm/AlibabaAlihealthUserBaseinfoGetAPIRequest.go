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

// New
