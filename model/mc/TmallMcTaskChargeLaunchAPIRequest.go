package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMcTaskChargeLaunchAPIRequest
云码充电宝投放链路 API请求
tmall.mc.task.charge.launch

云码充电宝投放链路，用于判断用户是否有匹配的投放计划 */
type TmallMcTaskChargeLaunchAPIRequest struct {
	model.Params
	// 外部设备编码
	_outerCode string
	// 渠道ID
	_channelId string
	// 支付宝openID
	_alipayOpenId string
	// urlID
	_urlId string
	// 服务商附加url参数
	_extra string
}

// New
