package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取商品列表 API请求
taobao.openmall.items.query

批量获取对联盟开放的商品列表。
*/
type TaobaoOpenmallItemsQueryRequest struct {
    model.Params
    // 已废弃，请勿使用
    _itemIds   string
    // 第几页，默认：1
    _pageNo   int64
    // 页大小，默认20，1~100
    _pageSize   int64
    // 当不输入渠道商时，展示全网公有商品池；当输入渠道商的淘宝Nick时，展示该渠道私有供给商品列表
    _distributor   string
}

// 初始化TaobaoOpenmallItemsQueryRequest对象
func NewTaobaoOpenmallItemsQueryRequest() *TaobaoOpenmallItemsQueryRequest{
    return &TaobaoOpenmallItemsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallItemsQueryRequest) GetApiMethodName() string {
    return "taobao.openmall.items.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallItemsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemIds Setter
// 已废弃，请勿使用
func (r *TaobaoOpenmallItemsQueryRequest) SetItemIds(_itemIds string) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoOpenmallItemsQueryRequest) GetItemIds() string {
    return r._itemIds
}
// PageNo Setter
// 第几页，默认：1
func (r *TaobaoOpenmallItemsQueryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoOpenmallItemsQueryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoOpenmallItemsQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpenmallItemsQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// Distributor Setter
// 当不输入渠道商时，展示全网公有商品池；当输入渠道商的淘宝Nick时，展示该渠道私有供给商品列表
func (r *TaobaoOpenmallItemsQueryRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallItemsQueryRequest) GetDistributor() string {
    return r._distributor
}
