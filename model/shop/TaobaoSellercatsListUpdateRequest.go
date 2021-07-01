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
type TaobaoSellercatsListUpdateAPIRequest struct {
    model.Params
    // 卖家自定义类目编号
    _cid   int64
    // 卖家自定义类目名称。不超过20个字符
    _name   string
    // 链接图片URL地址
    _pictUrl   string
    // 该类目在页面上的排序位置,取值范围:大于零的整数
    _sortOrder   int64
}

// 初始化TaobaoSellercatsListUpdateAPIRequest对象
func NewTaobaoSellercatsListUpdateRequest() *TaobaoSellercatsListUpdateAPIRequest{
    return &TaobaoSellercatsListUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSellercatsListUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.sellercats.list.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSellercatsListUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Cid Setter
// 卖家自定义类目编号
func (r *TaobaoSellercatsListUpdateAPIRequest) SetCid(_cid int64) error {
    r._cid = _cid
    r.Set("cid", _cid)
    return nil
}

// Cid Getter
func (r TaobaoSellercatsListUpdateAPIRequest) GetCid() int64 {
    return r._cid
}
// Name Setter
// 卖家自定义类目名称。不超过20个字符
func (r *TaobaoSellercatsListUpdateAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoSellercatsListUpdateAPIRequest) GetName() string {
    return r._name
}
// PictUrl Setter
// 链接图片URL地址
func (r *TaobaoSellercatsListUpdateAPIRequest) SetPictUrl(_pictUrl string) error {
    r._pictUrl = _pictUrl
    r.Set("pict_url", _pictUrl)
    return nil
}

// PictUrl Getter
func (r TaobaoSellercatsListUpdateAPIRequest) GetPictUrl() string {
    return r._pictUrl
}
// SortOrder Setter
// 该类目在页面上的排序位置,取值范围:大于零的整数
func (r *TaobaoSellercatsListUpdateAPIRequest) SetSortOrder(_sortOrder int64) error {
    r._sortOrder = _sortOrder
    r.Set("sort_order", _sortOrder)
    return nil
}

// SortOrder Getter
func (r TaobaoSellercatsListUpdateAPIRequest) GetSortOrder() int64 {
    return r._sortOrder
}
