package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCrmOmsRuleSyncAPIRequest
商家ERP订单处理规则同步 API请求
cainiao.crm.oms.rule.sync

将商家ERP订单处理规则同步到菜鸟CRM系统 */
type CainiaoCrmOmsRuleSyncAPIRequest struct {
	model.Params
	// 店铺nick
	_shopCode string
	// 是否开启菜鸟自动流转规则
	_isOpenCnauto bool
	// 是否系统智能处理订单（无人工介入）
	_isAutoCheck bool
	// 人工审单规则描述
	_checkRuleMsg string
	// 是否开启了订单合单
	_isSysMergeOrder bool
	// 订单合单周期（单位：分钟）
	_mergeOrderCycle int64
	// 其他未定义订单处理规则，格式｛name;stauts;cycle;｝
	_otherRule string
}

// New
