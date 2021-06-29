package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取商品列表 APIRequest
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

func NewTaobaoOpenmallItemsQueryRequest() *TaobaoOpenmallItemsQueryRequest{
    return &TaobaoOpenmallItemsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallItemsQueryRequest) GetApiMethodName() string {
    return "taobao.openmall.items.query"
}

func (r TaobaoOpenmallItemsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallItemsQueryRequest) SetItemIds(itemIds string) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoOpenmallItemsQueryRequest) GetItemIds() string {
    return r.itemIds
}

func (r *TaobaoOpenmallItemsQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoOpenmallItemsQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoOpenmallItemsQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoOpenmallItemsQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoOpenmallItemsQueryRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallItemsQueryRequest) GetDistributor() string {
    return r.distributor
}

