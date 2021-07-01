package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFmhealthButlerEnergySyncAPIRequest
同步用户消耗能量 API请求
alibaba.fmhealth.butler.energy.sync

同步用户消耗能量，用户消耗s点或卡路里后，同步给健康平台 */
type AlibabaFmhealthButlerEnergySyncAPIRequest struct {
	model.Params
	// 阿里用户id
	_userId int64
	// 每日已消耗能量
	_value *BigDecimal
	// “S”- s点 “CAL”- 卡路里
	_energyType string
	// 每日可消耗能量
	_target *BigDecimal
	// 每日运动消耗能量值
	_sport *BigDecimal
}

// New
