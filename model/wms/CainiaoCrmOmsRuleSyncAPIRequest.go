package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocrmomsrulesyncAPIRequest 商家ERP订单处理规则同步 API请求
// cainiao.crm.oms.rule.sync
//
// 将商家ERP订单处理规则同步到菜鸟CRM系统
type CainiaocrmomsrulesyncAPIRequest struct {
	model.Params
	// 店铺nick
	_shopCode string
	// 人工审单规则描述
	_checkRuleMsg string
	// 其他未定义订单处理规则，格式｛name;stauts;cycle;｝
	_otherRule string
	// 订单合单周期（单位：分钟）
	_mergeOrderCycle int64
	// 是否开启菜鸟自动流转规则
	_isOpenCnauto bool
	// 是否系统智能处理订单（无人工介入）
	_isAutoCheck bool
	// 是否开启了订单合单
	_isSysMergeOrder bool
}

// NewCainiaocrmomsrulesyncRequest 初始化CainiaocrmomsrulesyncAPIRequest对象
func NewCainiaocrmomsrulesyncRequest() *CainiaocrmomsrulesyncAPIRequest {
	return &CainiaocrmomsrulesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocrmomsrulesyncAPIRequest) GetApiMethodName() string {
	return "cainiao.crm.oms.rule.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocrmomsrulesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocrmomsrulesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopCode is ShopCode Setter
// 店铺nick
func (r *CainiaocrmomsrulesyncAPIRequest) SetShopCode(_shopCode string) error {
	r._shopCode = _shopCode
	r.Set("shop_code", _shopCode)
	return nil
}

// GetShopCode ShopCode Getter
func (r CainiaocrmomsrulesyncAPIRequest) GetShopCode() string {
	return r._shopCode
}

// SetCheckRuleMsg is CheckRuleMsg Setter
// 人工审单规则描述
func (r *CainiaocrmomsrulesyncAPIRequest) SetCheckRuleMsg(_checkRuleMsg string) error {
	r._checkRuleMsg = _checkRuleMsg
	r.Set("check_rule_msg", _checkRuleMsg)
	return nil
}

// GetCheckRuleMsg CheckRuleMsg Getter
func (r CainiaocrmomsrulesyncAPIRequest) GetCheckRuleMsg() string {
	return r._checkRuleMsg
}

// SetOtherRule is OtherRule Setter
// 其他未定义订单处理规则，格式｛name;stauts;cycle;｝
func (r *CainiaocrmomsrulesyncAPIRequest) SetOtherRule(_otherRule string) error {
	r._otherRule = _otherRule
	r.Set("other_rule", _otherRule)
	return nil
}

// GetOtherRule OtherRule Getter
func (r CainiaocrmomsrulesyncAPIRequest) GetOtherRule() string {
	return r._otherRule
}

// SetMergeOrderCycle is MergeOrderCycle Setter
// 订单合单周期（单位：分钟）
func (r *CainiaocrmomsrulesyncAPIRequest) SetMergeOrderCycle(_mergeOrderCycle int64) error {
	r._mergeOrderCycle = _mergeOrderCycle
	r.Set("merge_order_cycle", _mergeOrderCycle)
	return nil
}

// GetMergeOrderCycle MergeOrderCycle Getter
func (r CainiaocrmomsrulesyncAPIRequest) GetMergeOrderCycle() int64 {
	return r._mergeOrderCycle
}

// SetIsOpenCnauto is IsOpenCnauto Setter
// 是否开启菜鸟自动流转规则
func (r *CainiaocrmomsrulesyncAPIRequest) SetIsOpenCnauto(_isOpenCnauto bool) error {
	r._isOpenCnauto = _isOpenCnauto
	r.Set("is_open_cnauto", _isOpenCnauto)
	return nil
}

// GetIsOpenCnauto IsOpenCnauto Getter
func (r CainiaocrmomsrulesyncAPIRequest) GetIsOpenCnauto() bool {
	return r._isOpenCnauto
}

// SetIsAutoCheck is IsAutoCheck Setter
// 是否系统智能处理订单（无人工介入）
func (r *CainiaocrmomsrulesyncAPIRequest) SetIsAutoCheck(_isAutoCheck bool) error {
	r._isAutoCheck = _isAutoCheck
	r.Set("is_auto_check", _isAutoCheck)
	return nil
}

// GetIsAutoCheck IsAutoCheck Getter
func (r CainiaocrmomsrulesyncAPIRequest) GetIsAutoCheck() bool {
	return r._isAutoCheck
}

// SetIsSysMergeOrder is IsSysMergeOrder Setter
// 是否开启了订单合单
func (r *CainiaocrmomsrulesyncAPIRequest) SetIsSysMergeOrder(_isSysMergeOrder bool) error {
	r._isSysMergeOrder = _isSysMergeOrder
	r.Set("is_sys_merge_order", _isSysMergeOrder)
	return nil
}

// GetIsSysMergeOrder IsSysMergeOrder Getter
func (r CainiaocrmomsrulesyncAPIRequest) GetIsSysMergeOrder() bool {
	return r._isSysMergeOrder
}
