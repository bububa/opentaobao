package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunUnimktTaskChargeLaunchAPIRequest
云码权益查询 API请求
aliyun.unimkt.task.charge.launch

云码线上流量投放链路，用于判断用户是否有匹配的投放计划 */
type AliyunUnimktTaskChargeLaunchAPIRequest struct {
	model.Params
	// 服务商附加url参数
	_extra string
	// urlID
	_urlId string
	// 支付宝openID
	_alipayOpenId string
	// 渠道ID
	_channelId string
	// 淘宝ID
	_userId string
}

// New
