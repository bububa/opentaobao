package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家ERP订单处理规则同步 API请求
cainiao.crm.oms.rule.sync

将商家ERP订单处理规则同步到菜鸟CRM系统
*/
type CainiaoCrmOmsRuleSyncRequest struct {
    model.Params
    // 店铺nick
    _shopCode   string
    // 是否开启菜鸟自动流转规则
    _isOpenCnauto   bool
    // 是否系统智能处理订单（无人工介入）
    _isAutoCheck   bool
    // 人工审单规则描述
    _checkRuleMsg   string
    // 是否开启了订单合单
    _isSysMergeOrder   bool
    // 订单合单周期（单位：分钟）
    _mergeOrderCycle   int64
    // 其他未定义订单处理规则，格式｛name;stauts;cycle;｝
    _otherRule   string
}

// 初始化CainiaoCrmOmsRuleSyncRequest对象
func NewCainiaoCrmOmsRuleSyncRequest() *CainiaoCrmOmsRuleSyncRequest{
    return &CainiaoCrmOmsRuleSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCrmOmsRuleSyncRequest) GetApiMethodName() string {
    return "cainiao.crm.oms.rule.sync"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCrmOmsRuleSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopCode Setter
// 店铺nick
func (r *CainiaoCrmOmsRuleSyncRequest) SetShopCode(_shopCode string) error {
    r._shopCode = _shopCode
    r.Set("shop_code", _shopCode)
    return nil
}

// ShopCode Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetShopCode() string {
    return r._shopCode
}
// IsOpenCnauto Setter
// 是否开启菜鸟自动流转规则
func (r *CainiaoCrmOmsRuleSyncRequest) SetIsOpenCnauto(_isOpenCnauto bool) error {
    r._isOpenCnauto = _isOpenCnauto
    r.Set("is_open_cnauto", _isOpenCnauto)
    return nil
}

// IsOpenCnauto Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetIsOpenCnauto() bool {
    return r._isOpenCnauto
}
// IsAutoCheck Setter
// 是否系统智能处理订单（无人工介入）
func (r *CainiaoCrmOmsRuleSyncRequest) SetIsAutoCheck(_isAutoCheck bool) error {
    r._isAutoCheck = _isAutoCheck
    r.Set("is_auto_check", _isAutoCheck)
    return nil
}

// IsAutoCheck Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetIsAutoCheck() bool {
    return r._isAutoCheck
}
// CheckRuleMsg Setter
// 人工审单规则描述
func (r *CainiaoCrmOmsRuleSyncRequest) SetCheckRuleMsg(_checkRuleMsg string) error {
    r._checkRuleMsg = _checkRuleMsg
    r.Set("check_rule_msg", _checkRuleMsg)
    return nil
}

// CheckRuleMsg Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetCheckRuleMsg() string {
    return r._checkRuleMsg
}
// IsSysMergeOrder Setter
// 是否开启了订单合单
func (r *CainiaoCrmOmsRuleSyncRequest) SetIsSysMergeOrder(_isSysMergeOrder bool) error {
    r._isSysMergeOrder = _isSysMergeOrder
    r.Set("is_sys_merge_order", _isSysMergeOrder)
    return nil
}

// IsSysMergeOrder Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetIsSysMergeOrder() bool {
    return r._isSysMergeOrder
}
// MergeOrderCycle Setter
// 订单合单周期（单位：分钟）
func (r *CainiaoCrmOmsRuleSyncRequest) SetMergeOrderCycle(_mergeOrderCycle int64) error {
    r._mergeOrderCycle = _mergeOrderCycle
    r.Set("merge_order_cycle", _mergeOrderCycle)
    return nil
}

// MergeOrderCycle Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetMergeOrderCycle() int64 {
    return r._mergeOrderCycle
}
// OtherRule Setter
// 其他未定义订单处理规则，格式｛name;stauts;cycle;｝
func (r *CainiaoCrmOmsRuleSyncRequest) SetOtherRule(_otherRule string) error {
    r._otherRule = _otherRule
    r.Set("other_rule", _otherRule)
    return nil
}

// OtherRule Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetOtherRule() string {
    return r._otherRule
}
