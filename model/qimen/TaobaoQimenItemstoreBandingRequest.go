package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品关联绑定接口 API请求
taobao.qimen.itemstore.banding

商家在ERP等系统中调用该接口，将线上商品和线下门店“新建/删除”关联。这里的线上。每次只能单个商品关联多个门店，门店上限200
*/
type TaobaoQimenItemstoreBandingRequest struct {
    model.Params
    // 门店列表
    storeIds   []int64
    // 备注信息
    remark   string
    // 操作类型
    actionType   string
    // 线上商品ID
    itemId   int64
}

// 初始化TaobaoQimenItemstoreBandingRequest对象
func NewTaobaoQimenItemstoreBandingRequest() *TaobaoQimenItemstoreBandingRequest{
    return &TaobaoQimenItemstoreBandingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemstoreBandingRequest) GetApiMethodName() string {
    return "taobao.qimen.itemstore.banding"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemstoreBandingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreIds Setter
// 门店列表
func (r *TaobaoQimenItemstoreBandingRequest) SetStoreIds(storeIds []int64) error {
    r.storeIds = storeIds
    r.Set("store_ids", storeIds)
    return nil
}

// StoreIds Getter
func (r TaobaoQimenItemstoreBandingRequest) GetStoreIds() []int64 {
    return r.storeIds
}
// Remark Setter
// 备注信息
func (r *TaobaoQimenItemstoreBandingRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoQimenItemstoreBandingRequest) GetRemark() string {
    return r.remark
}
// ActionType Setter
// 操作类型
func (r *TaobaoQimenItemstoreBandingRequest) SetActionType(actionType string) error {
    r.actionType = actionType
    r.Set("action_type", actionType)
    return nil
}

// ActionType Getter
func (r TaobaoQimenItemstoreBandingRequest) GetActionType() string {
    return r.actionType
}
// ItemId Setter
// 线上商品ID
func (r *TaobaoQimenItemstoreBandingRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoQimenItemstoreBandingRequest) GetItemId() int64 {
    return r.itemId
}
