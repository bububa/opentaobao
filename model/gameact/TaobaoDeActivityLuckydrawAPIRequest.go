package gameact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDeActivityLuckydrawAPIRequest
抽奖 API请求
taobao.de.activity.luckydraw

用于激励平台对外提供抽奖功能，包括但不限于集分宝、红包、宝点、淘金币、淘彩票等 */
type TaobaoDeActivityLuckydrawAPIRequest struct {
	model.Params
	// 运营和cp约定的事件唯一标示
	_eventKey string
	// 时间戳
	_sequenceId int64
	// 用户的串ID
	_accountId string
	// 机器设备号
	_machineId string
	// 确认签名key
	_confirmKey string
	// 行为Key
	_behaviorKey string
	// 渠道
	_channel string
	// 使用市场
	_market string
	// 盒型号
	_deviceModel string
	// 魔盒分发渠道
	_distribChannel string
	// 魔盒UUID
	_uuid string
}

// New
