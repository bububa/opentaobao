package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品关联绑定接口 APIRequest
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

func NewTaobaoQimenItemstoreBandingRequest() *TaobaoQimenItemstoreBandingRequest{
    return &TaobaoQimenItemstoreBandingRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenItemstoreBandingRequest) GetApiMethodName() string {
    return "taobao.qimen.itemstore.banding"
}

func (r TaobaoQimenItemstoreBandingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenItemstoreBandingRequest) SetStoreIds(storeIds []int64) error {
    r.storeIds = storeIds
    r.Set("store_ids", storeIds)
    return nil
}

func (r TaobaoQimenItemstoreBandingRequest) GetStoreIds() []int64 {
    return r.storeIds
}

func (r *TaobaoQimenItemstoreBandingRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoQimenItemstoreBandingRequest) GetRemark() string {
    return r.remark
}

func (r *TaobaoQimenItemstoreBandingRequest) SetActionType(actionType string) error {
    r.actionType = actionType
    r.Set("action_type", actionType)
    return nil
}

func (r TaobaoQimenItemstoreBandingRequest) GetActionType() string {
    return r.actionType
}

func (r *TaobaoQimenItemstoreBandingRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoQimenItemstoreBandingRequest) GetItemId() int64 {
    return r.itemId
}

