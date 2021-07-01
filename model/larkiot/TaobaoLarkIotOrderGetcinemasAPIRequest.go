package larkiot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLarkIotOrderGetcinemasAPIRequest
获取iot渠道开放的影院 API请求
taobao.lark.iot.order.getcinemas

iot渠道拉取有权限访问的影院 */
type TaobaoLarkIotOrderGetcinemasAPIRequest struct {
	model.Params
	// 渠道编码
	_channelCode string
}

// New
