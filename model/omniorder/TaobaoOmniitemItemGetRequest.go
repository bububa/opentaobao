package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全渠道门店商品 APIRequest
taobao.omniitem.item.get

通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息
*/
type TaobaoOmniitemItemGetRequest struct {
    model.Params

    // 分页当前页数
    pageNo   int64 

    // 分页单页大小
    pageSize   int64 

    // 可选，指定获取的商品id
    itemId   int64 

    // 可选，指定获取的商品外部id
    outerId   string 

}

func NewTaobaoOmniitemItemGetRequest() *TaobaoOmniitemItemGetRequest{
    return &TaobaoOmniitemItemGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemItemGetRequest) GetApiMethodName() string {
    return "taobao.omniitem.item.get"
}

func (r TaobaoOmniitemItemGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemItemGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoOmniitemItemGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoOmniitemItemGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoOmniitemItemGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoOmniitemItemGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOmniitemItemGetRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoOmniitemItemGetRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoOmniitemItemGetRequest) GetOuterId() string {
    return r.outerId
}

