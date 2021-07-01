package wenyuvideo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuWenyuvideoPersionGetAPIRequest
根据优酷人物ID获取人物详情页，包含相关影视和相关人物 API请求
youku.wenyuvideo.persion.get

根据优酷人物ID获取人物详情页，包含相关影视和相关人物 */
type YoukuWenyuvideoPersionGetAPIRequest struct {
	model.Params
	// 设备信息
	_systemInfo string
	// 人物ID
	_personId int64
}

// New
