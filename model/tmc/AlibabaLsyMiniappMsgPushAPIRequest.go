package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyMiniappMsgPushAPIRequest
零售云小程序消息推送 API请求
alibaba.lsy.miniapp.msg.push

零售云小程序消息推送，推送消息至零售云（喵零等） */
type AlibabaLsyMiniappMsgPushAPIRequest struct {
	model.Params
	// 小程序ID
	_appId string
	// 消息ID
	_msgId int64
	// 摊位ID
	_storeId int64
	// 消息模板，miaoling_msg_isv_clue - 线索通知消息
	_templateId string
	// 消息参数
	_params string
}

// New
