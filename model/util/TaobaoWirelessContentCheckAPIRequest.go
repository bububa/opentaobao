package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWirelessContentCheckAPIRequest
无线开放内容检查 API请求
taobao.wireless.content.check

无线开放内容检查，提供涉黄暴力政治文本检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70439.html" target="blank">阿里云内容安全</a> */
type TaobaoWirelessContentCheckAPIRequest struct {
	model.Params
	// 待检查的文本
	_text string
}

// New
