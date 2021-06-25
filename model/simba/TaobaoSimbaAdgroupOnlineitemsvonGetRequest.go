package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户上架在线销售的全部宝贝 APIRequest
taobao.simba.adgroup.onlineitemsvon.get

获取用户上架在线销售的全部宝贝
*/
type TaobaoSimbaAdgroupOnlineitemsvonGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 排序字段，starts：按开始时间排序bidCount:按销量排序
    orderField   string 

    // 排序，true:降序， false:升序
    orderBy   bool 

    // 页尺寸，最大200
    pageSize   int64 

    // 页码，从1开始,最大50。最大只能获取1W个宝贝
    pageNo   int64 

    // 推广单元类型 101001005代表标准推广，101001014代表销量明星推广
    productId   int64 

}

func NewTaobaoSimbaAdgroupOnlineitemsvonGetRequest() *TaobaoSimbaAdgroupOnlineitemsvonGetRequest{
    return &TaobaoSimbaAdgroupOnlineitemsvonGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetApiMethodName() string {
    return "taobao.simba.adgroup.onlineitemsvon.get"
}

func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetOrderField(orderField string) error {
    r.orderField = orderField
    r.Set("order_field", orderField)
    return nil
}

func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetOrderField() string {
    return r.orderField
}

func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetOrderBy(orderBy bool) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetOrderBy() bool {
    return r.orderBy
}

func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetProductId() int64 {
    return r.productId
}

