package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMessageAddtextAPIRequest
精灵代说 API请求
taobao.ailab.aicloud.top.message.addtext

精灵代说 */
type TaobaoAilabAicloudTopMessageAddtextAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 代说文本
	_param2 string
	// 扩展信息，可以为空
	_param3 string
}

// New
