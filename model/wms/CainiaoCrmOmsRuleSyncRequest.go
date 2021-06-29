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
    shopCode   string
    // 是否开启菜鸟自动流转规则
    isOpenCnauto   bool
    // 是否系统智能处理订单（无人工介入）
    isAutoCheck   bool
    // 人工审单规则描述
    checkRuleMsg   string
    // 是否开启了订单合单
    isSysMergeOrder   bool
    // 订单合单周期（单位：分钟）
    mergeOrderCycle   int64
    // 其他未定义订单处理规则，格式｛name;stauts;cycle;｝
    otherRule   string
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
func (r *CainiaoCrmOmsRuleSyncRequest) SetShopCode(shopCode string) error {
    r.shopCode = shopCode
    r.Set("shop_code", shopCode)
    return nil
}

// ShopCode Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetShopCode() string {
    return r.shopCode
}
// IsOpenCnauto Setter
// 是否开启菜鸟自动流转规则
func (r *CainiaoCrmOmsRuleSyncRequest) SetIsOpenCnauto(isOpenCnauto bool) error {
    r.isOpenCnauto = isOpenCnauto
    r.Set("is_open_cnauto", isOpenCnauto)
    return nil
}

// IsOpenCnauto Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetIsOpenCnauto() bool {
    return r.isOpenCnauto
}
// IsAutoCheck Setter
// 是否系统智能处理订单（无人工介入）
func (r *CainiaoCrmOmsRuleSyncRequest) SetIsAutoCheck(isAutoCheck bool) error {
    r.isAutoCheck = isAutoCheck
    r.Set("is_auto_check", isAutoCheck)
    return nil
}

// IsAutoCheck Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetIsAutoCheck() bool {
    return r.isAutoCheck
}
// CheckRuleMsg Setter
// 人工审单规则描述
func (r *CainiaoCrmOmsRuleSyncRequest) SetCheckRuleMsg(checkRuleMsg string) error {
    r.checkRuleMsg = checkRuleMsg
    r.Set("check_rule_msg", checkRuleMsg)
    return nil
}

// CheckRuleMsg Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetCheckRuleMsg() string {
    return r.checkRuleMsg
}
// IsSysMergeOrder Setter
// 是否开启了订单合单
func (r *CainiaoCrmOmsRuleSyncRequest) SetIsSysMergeOrder(isSysMergeOrder bool) error {
    r.isSysMergeOrder = isSysMergeOrder
    r.Set("is_sys_merge_order", isSysMergeOrder)
    return nil
}

// IsSysMergeOrder Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetIsSysMergeOrder() bool {
    return r.isSysMergeOrder
}
// MergeOrderCycle Setter
// 订单合单周期（单位：分钟）
func (r *CainiaoCrmOmsRuleSyncRequest) SetMergeOrderCycle(mergeOrderCycle int64) error {
    r.mergeOrderCycle = mergeOrderCycle
    r.Set("merge_order_cycle", mergeOrderCycle)
    return nil
}

// MergeOrderCycle Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetMergeOrderCycle() int64 {
    return r.mergeOrderCycle
}
// OtherRule Setter
// 其他未定义订单处理规则，格式｛name;stauts;cycle;｝
func (r *CainiaoCrmOmsRuleSyncRequest) SetOtherRule(otherRule string) error {
    r.otherRule = otherRule
    r.Set("other_rule", otherRule)
    return nil
}

// OtherRule Getter
func (r CainiaoCrmOmsRuleSyncRequest) GetOtherRule() string {
    return r.otherRule
}
