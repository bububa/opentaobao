package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotUseroscontrolAPIRequest
物联卡用户管理停开机功能 API请求
alibaba.aliqin.fc.iot.useroscontrol

物联网针对用户级管理停开支持 */
type AlibabaAliqinFcIotUseroscontrolAPIRequest struct {
	model.Params
	// 物联卡的iccid
	_iccid string
	// 用户停开的操作类型：MANAGE_RESUME、MANAGE_STOP
	_action string
}

// New
