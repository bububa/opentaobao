package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新卖家自定义类目 API请求
taobao.sellercats.list.update

此API更新卖家店铺内自定义类目 <br/>注：因为缓存的关系，添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
*/
type TaobaoSellercatsListUpdateRequest struct {
    model.Params
    // 卖家自定义类目编号
    cid   int64
    // 卖家自定义类目名称。不超过20个字符
    name   string
    // 链接图片URL地址
    pictUrl   string
    // 该类目在页面上的排序位置,取值范围:大于零的整数
    sortOrder   int64
}

// 初始化TaobaoSellercatsListUpdateRequest对象
func NewTaobaoSellercatsListUpdateRequest() *TaobaoSellercatsListUpdateRequest{
    return &TaobaoSellercatsListUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSellercatsListUpdateRequest) GetApiMethodName() string {
    return "taobao.sellercats.list.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSellercatsListUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Cid Setter
// 卖家自定义类目编号
func (r *TaobaoSellercatsListUpdateRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

// Cid Getter
func (r TaobaoSellercatsListUpdateRequest) GetCid() int64 {
    return r.cid
}
// Name Setter
// 卖家自定义类目名称。不超过20个字符
func (r *TaobaoSellercatsListUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoSellercatsListUpdateRequest) GetName() string {
    return r.name
}
// PictUrl Setter
// 链接图片URL地址
func (r *TaobaoSellercatsListUpdateRequest) SetPictUrl(pictUrl string) error {
    r.pictUrl = pictUrl
    r.Set("pict_url", pictUrl)
    return nil
}

// PictUrl Getter
func (r TaobaoSellercatsListUpdateRequest) GetPictUrl() string {
    return r.pictUrl
}
// SortOrder Setter
// 该类目在页面上的排序位置,取值范围:大于零的整数
func (r *TaobaoSellercatsListUpdateRequest) SetSortOrder(sortOrder int64) error {
    r.sortOrder = sortOrder
    r.Set("sort_order", sortOrder)
    return nil
}

// SortOrder Getter
func (r TaobaoSellercatsListUpdateRequest) GetSortOrder() int64 {
    return r.sortOrder
}
