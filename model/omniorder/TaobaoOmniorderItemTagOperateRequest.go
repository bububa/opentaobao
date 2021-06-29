package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品打标与去标 API请求
taobao.omniorder.item.tag.operate

用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。
*/
type TaobaoOmniorderItemTagOperateRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
    // 商品标,storeDeliver代表门店发货, AllocateByFront代表前置拆单, storeCollect代表门店自提
    _types   []string
    // 操作状态， 填 1 代表打标，填 -1 代表去标
    _status   int64
    // 分单&接单设置
    _omniSetting   *OmniSettingDTO
}

// 初始化TaobaoOmniorderItemTagOperateRequest对象
func NewTaobaoOmniorderItemTagOperateRequest() *TaobaoOmniorderItemTagOperateRequest{
    return &TaobaoOmniorderItemTagOperateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderItemTagOperateRequest) GetApiMethodName() string {
    return "taobao.omniorder.item.tag.operate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderItemTagOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TaobaoOmniorderItemTagOperateRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOmniorderItemTagOperateRequest) GetItemId() int64 {
    return r._itemId
}
// Types Setter
// 商品标,storeDeliver代表门店发货, AllocateByFront代表前置拆单, storeCollect代表门店自提
func (r *TaobaoOmniorderItemTagOperateRequest) SetTypes(_types []string) error {
    r._types = _types
    r.Set("types", _types)
    return nil
}

// Types Getter
func (r TaobaoOmniorderItemTagOperateRequest) GetTypes() []string {
    return r._types
}
// Status Setter
// 操作状态， 填 1 代表打标，填 -1 代表去标
func (r *TaobaoOmniorderItemTagOperateRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoOmniorderItemTagOperateRequest) GetStatus() int64 {
    return r._status
}
// OmniSetting Setter
// 分单&接单设置
func (r *TaobaoOmniorderItemTagOperateRequest) SetOmniSetting(_omniSetting *OmniSettingDTO) error {
    r._omniSetting = _omniSetting
    r.Set("omni_setting", _omniSetting)
    return nil
}

// OmniSetting Getter
func (r TaobaoOmniorderItemTagOperateRequest) GetOmniSetting() *OmniSettingDTO {
    return r._omniSetting
}
