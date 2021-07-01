package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest
切换用户 API请求
alibaba.ailabs.tmallgenie.auth.switchuser

设备切换授权用户 */
type AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest struct {
	model.Params
	// client_id
	_clientId string
	// 目标用户openId
	_newUserOpenId string
	// 当前拥有设备权限的用户openId
	_oldUserOpenId string
	// 设备uuid
	_uuid string
}

// New
