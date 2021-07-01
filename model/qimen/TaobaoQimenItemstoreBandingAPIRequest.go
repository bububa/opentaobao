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
type TaobaoQimenItemstoreBandingAPIRequest struct {
    model.Params
    // 门店列表
    _storeIds   []int64
    // 备注信息
    _remark   string
    // 操作类型
    _actionType   string
    // 线上商品ID
    _itemId   int64
}

// 初始化TaobaoQimenItemstoreBandingAPIRequest对象
func NewTaobaoQimenItemstoreBandingRequest() *TaobaoQimenItemstoreBandingAPIRequest{
    return &TaobaoQimenItemstoreBandingAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemstoreBandingAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.itemstore.banding"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemstoreBandingAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreIds Setter
// 门店列表
func (r *TaobaoQimenItemstoreBandingAPIRequest) SetStoreIds(_storeIds []int64) error {
    r._storeIds = _storeIds
    r.Set("store_ids", _storeIds)
    return nil
}

// StoreIds Getter
func (r TaobaoQimenItemstoreBandingAPIRequest) GetStoreIds() []int64 {
    return r._storeIds
}
// Remark Setter
// 备注信息
func (r *TaobaoQimenItemstoreBandingAPIRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoQimenItemstoreBandingAPIRequest) GetRemark() string {
    return r._remark
}
// ActionType Setter
// 操作类型
func (r *TaobaoQimenItemstoreBandingAPIRequest) SetActionType(_actionType string) error {
    r._actionType = _actionType
    r.Set("action_type", _actionType)
    return nil
}

// ActionType Getter
func (r TaobaoQimenItemstoreBandingAPIRequest) GetActionType() string {
    return r._actionType
}
// ItemId Setter
// 线上商品ID
func (r *TaobaoQimenItemstoreBandingAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoQimenItemstoreBandingAPIRequest) GetItemId() int64 {
    return r._itemId
}
