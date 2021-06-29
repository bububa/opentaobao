package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品通自动打标 API请求
taobao.qimen.items.marking

调用该接口，对商品进行XXXX标的打标、去标的动作。
*/
type TaobaoQimenItemsMarkingRequest struct {
    model.Params
    // 操作类型，string（50），ADD=打标，DELETE=去标，必填
    actionType   string
    // 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
    tagType   string
    // 线上商品ID，long，必填
    itemIds   []int64
    // 备注，string（500）
    remark   string
}

// 初始化TaobaoQimenItemsMarkingRequest对象
func NewTaobaoQimenItemsMarkingRequest() *TaobaoQimenItemsMarkingRequest{
    return &TaobaoQimenItemsMarkingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemsMarkingRequest) GetApiMethodName() string {
    return "taobao.qimen.items.marking"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemsMarkingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActionType Setter
// 操作类型，string（50），ADD=打标，DELETE=去标，必填
func (r *TaobaoQimenItemsMarkingRequest) SetActionType(actionType string) error {
    r.actionType = actionType
    r.Set("action_type", actionType)
    return nil
}

// ActionType Getter
func (r TaobaoQimenItemsMarkingRequest) GetActionType() string {
    return r.actionType
}
// TagType Setter
// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
func (r *TaobaoQimenItemsMarkingRequest) SetTagType(tagType string) error {
    r.tagType = tagType
    r.Set("tag_type", tagType)
    return nil
}

// TagType Getter
func (r TaobaoQimenItemsMarkingRequest) GetTagType() string {
    return r.tagType
}
// ItemIds Setter
// 线上商品ID，long，必填
func (r *TaobaoQimenItemsMarkingRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoQimenItemsMarkingRequest) GetItemIds() []int64 {
    return r.itemIds
}
// Remark Setter
// 备注，string（500）
func (r *TaobaoQimenItemsMarkingRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoQimenItemsMarkingRequest) GetRemark() string {
    return r.remark
}
