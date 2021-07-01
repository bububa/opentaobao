package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcApRuleCreateAPIRequest
创建分账规则 API请求
taobao.oc.ap.rule.create

OC分账业务功能支持，用于创建分账规则 */
type TaobaoOcApRuleCreateAPIRequest struct {
	model.Params
	// 传入比例数组后者金额数组
	_divisionRule []int64
	// 规则描述相关扩展信息，divisonRule的值包含（"byAmount" 或者 "byPercentage"），excutionPeriod的值包含（ "month" 或者 "day" 或者 "now"）
	_extAttributes string
}

// New
