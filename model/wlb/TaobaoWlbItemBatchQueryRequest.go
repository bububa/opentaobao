package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批次库存查询接口 API请求
taobao.wlb.item.batch.query

根据用户id，item id list和store code来查询商品库存信息和批次信息
*/
type TaobaoWlbItemBatchQueryAPIRequest struct {
    model.Params
    // 仓库编号
    _storeCode   string
    // 分页查询参数，指定查询页数，默认为1
    _pageNo   int64
    // 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询
    _pageSize   int64
    // 需要查询的商品ID列表，以字符串表示，ID间以;隔开
    _itemIds   string
}

// 初始化TaobaoWlbItemBatchQueryAPIRequest对象
func NewTaobaoWlbItemBatchQueryRequest() *TaobaoWlbItemBatchQueryAPIRequest{
    return &TaobaoWlbItemBatchQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemBatchQueryAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.item.batch.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemBatchQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// 仓库编号
func (r *TaobaoWlbItemBatchQueryAPIRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbItemBatchQueryAPIRequest) GetStoreCode() string {
    return r._storeCode
}
// PageNo Setter
// 分页查询参数，指定查询页数，默认为1
func (r *TaobaoWlbItemBatchQueryAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbItemBatchQueryAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询
func (r *TaobaoWlbItemBatchQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbItemBatchQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// ItemIds Setter
// 需要查询的商品ID列表，以字符串表示，ID间以;隔开
func (r *TaobaoWlbItemBatchQueryAPIRequest) SetItemIds(_itemIds string) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoWlbItemBatchQueryAPIRequest) GetItemIds() string {
    return r._itemIds
}
