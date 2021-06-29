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
    _nick   string
    // 排序字段，starts：按开始时间排序bidCount:按销量排序
    _orderField   string
    // 排序，true:降序， false:升序
    _orderBy   bool
    // 页尺寸，最大200
    _pageSize   int64
    // 页码，从1开始,最大50。最大只能获取1W个宝贝
    _pageNo   int64
    // 推广单元类型 101001005代表标准推广，101001014代表销量明星推广
    _productId   int64
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
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetNick() string {
    return r._nick
}
// OrderField Setter
// 排序字段，starts：按开始时间排序bidCount:按销量排序
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetOrderField(_orderField string) error {
    r._orderField = _orderField
    r.Set("order_field", _orderField)
    return nil
}

// OrderField Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetOrderField() string {
    return r._orderField
}
// OrderBy Setter
// 排序，true:降序， false:升序
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetOrderBy(_orderBy bool) error {
    r._orderBy = _orderBy
    r.Set("order_by", _orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetOrderBy() bool {
    return r._orderBy
}
// PageSize Setter
// 页尺寸，最大200
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 页码，从1开始,最大50。最大只能获取1W个宝贝
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// ProductId Setter
// 推广单元类型 101001005代表标准推广，101001014代表销量明星推广
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetRequest) GetProductId() int64 {
    return r._productId
}
