package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopFeedlistGetAPIRequest
获取对话流列表 API请求
taobao.ailab.aicloud.top.feedlist.get

获取指定应用的对话流信息 */
type TaobaoAilabAicloudTopFeedlistGetAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 最后一条对话的key
	_param2 string
	// 单页的条目数，注意，是String类型！
	_param3 string
}

// New
