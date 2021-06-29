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
    itemIds   string
    // 第几页，默认：1
    pageNo   int64
    // 页大小，默认20，1~100
    pageSize   int64
    // 当不输入渠道商时，展示全网公有商品池；当输入渠道商的淘宝Nick时，展示该渠道私有供给商品列表
    distributor   string
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
func (r *TaobaoOpenmallItemsQueryRequest) SetItemIds(itemIds string) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoOpenmallItemsQueryRequest) GetItemIds() string {
    return r.itemIds
}
// PageNo Setter
// 第几页，默认：1
func (r *TaobaoOpenmallItemsQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoOpenmallItemsQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoOpenmallItemsQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpenmallItemsQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// Distributor Setter
// 当不输入渠道商时，展示全网公有商品池；当输入渠道商的淘宝Nick时，展示该渠道私有供给商品列表
func (r *TaobaoOpenmallItemsQueryRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallItemsQueryRequest) GetDistributor() string {
    return r.distributor
}
