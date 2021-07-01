package lbs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLbsMessageUploadAPIRequest
lbs数据采集 API请求
taobao.lbs.message.upload

lbs数据采集 */
type TaobaoLbsMessageUploadAPIRequest struct {
	model.Params
	// 消息TOPIC
	_topic string
	// 消息体 json结构
	_body string
}

// New
