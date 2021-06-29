package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户上架在线销售的全部宝贝 API请求
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

// 初始化TaobaoSimbaAdgroupOnlineitemsvonGetRequest对象
func NewTaobaoSimbaAdgroupOnlineitemsvonGetRequest() *TaobaoSimbaAdgroupOnlineitemsvonGetRequest{
    return &TaobaoSimbaAdgroupOnlineitemsvonGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetApiMethodName() string {
    return "taobao.simba.adgroup.onlineitemsvon.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetNick() string {
    return r.nick
}
// OrderField Setter
// 排序字段，starts：按开始时间排序bidCount:按销量排序
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetOrderField(orderField string) error {
    r.orderField = orderField
    r.Set("order_field", orderField)
    return nil
}

// OrderField Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetOrderField() string {
    return r.orderField
}
// OrderBy Setter
// 排序，true:降序， false:升序
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetOrderBy(orderBy bool) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetOrderBy() bool {
    return r.orderBy
}
// PageSize Setter
// 页尺寸，最大200
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNo Setter
// 页码，从1开始,最大50。最大只能获取1W个宝贝
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// ProductId Setter
// 推广单元类型 101001005代表标准推广，101001014代表销量明星推广
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetProductId() int64 {
    return r.productId
}
