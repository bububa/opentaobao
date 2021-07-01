package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsLdlserverWxeventPostAPIRequest
乐动力线-微信设备绑定通知 API请求
alibaba.alisports.ldlserver.wxevent.post

转发弹外微信服务到弹内 */
type AlibabaAlisportsLdlserverWxeventPostAPIRequest struct {
	model.Params
	// 微信服务器消息数据
	_rawstr string
}

// New
