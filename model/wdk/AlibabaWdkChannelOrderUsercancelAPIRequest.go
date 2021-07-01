package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkChannelOrderUsercancelAPIRequest
用户发起售中取消 API请求
alibaba.wdk.channel.order.usercancel

用户发起售中取消 */
type AlibabaWdkChannelOrderUsercancelAPIRequest struct {
	model.Params
	// 取消信息
	_userCancelInfo *OrderUserCancelInfo
}

// New
