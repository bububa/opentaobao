package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品打标与去标 APIRequest
taobao.omniorder.item.tag.operate

用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。
*/
type TaobaoOmniorderItemTagOperateRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

    // 商品标,storeDeliver代表门店发货, AllocateByFront代表前置拆单, storeCollect代表门店自提
    types   []string 

    // 操作状态， 填 1 代表打标，填 -1 代表去标
    status   int64 

    // 分单&接单设置
    omniSetting   *OmniSettingDto 

}

func NewTaobaoOmniorderItemTagOperateRequest() *TaobaoOmniorderItemTagOperateRequest{
    return &TaobaoOmniorderItemTagOperateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderItemTagOperateRequest) GetApiMethodName() string {
    return "taobao.omniorder.item.tag.operate"
}

func (r TaobaoOmniorderItemTagOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderItemTagOperateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOmniorderItemTagOperateRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoOmniorderItemTagOperateRequest) SetTypes(types []string) error {
    r.types = types
    r.Set("types", types)
    return nil
}

func (r TaobaoOmniorderItemTagOperateRequest) GetTypes() []string {
    return r.types
}

func (r *TaobaoOmniorderItemTagOperateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoOmniorderItemTagOperateRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoOmniorderItemTagOperateRequest) SetOmniSetting(omniSetting *OmniSettingDto) error {
    r.omniSetting = omniSetting
    r.Set("omni_setting", omniSetting)
    return nil
}

func (r TaobaoOmniorderItemTagOperateRequest) GetOmniSetting() *OmniSettingDto {
    return r.omniSetting
}

