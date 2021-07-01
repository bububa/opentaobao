package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowPublishAPIRequest
流量发放(用户id) API请求
alibaba.aliqin.flow.publish

阿里通信流量下发功能，允许用户补发 */
type AlibabaAliqinFlowPublishAPIRequest struct {
	model.Params
	// 用户id
	_userId string
	// 流量
	_flow string
	// 原因
	_reason string
	// 唯一流水号（字母+数字）
	_serial string
	// 设置true为始终发送成功
	_always string
}

// New
