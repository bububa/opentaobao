package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加卖家自定义类目 API请求
taobao.sellercats.list.add

此API添加卖家店铺内自定义类目 <br/>父类目parent_cid值等于0：表示此类目为店铺下的一级类目，值不等于0：表示此类目有父类目 <br/>注：因为缓存的关系,添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
*/
type TaobaoSellercatsListAddRequest struct {
    model.Params
    // 卖家自定义类目名称。不超过20个字符
    _name   string
    // 链接图片URL地址。(绝对地址，格式：http://host/image_path)
    _pictUrl   string
    // 父类目编号，如果类目为店铺下的一级类目：值等于0，如果类目为子类目，调用获取taobao.sellercats.list.get父类目编号
    _parentCid   int64
    // 该类目在页面上的排序位置,取值范围:大于零的整数
    _sortOrder   int64
}

// 初始化TaobaoSellercatsListAddRequest对象
func NewTaobaoSellercatsListAddRequest() *TaobaoSellercatsListAddRequest{
    return &TaobaoSellercatsListAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSellercatsListAddRequest) GetApiMethodName() string {
    return "taobao.sellercats.list.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSellercatsListAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 卖家自定义类目名称。不超过20个字符
func (r *TaobaoSellercatsListAddRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoSellercatsListAddRequest) GetName() string {
    return r._name
}
// PictUrl Setter
// 链接图片URL地址。(绝对地址，格式：http://host/image_path)
func (r *TaobaoSellercatsListAddRequest) SetPictUrl(_pictUrl string) error {
    r._pictUrl = _pictUrl
    r.Set("pict_url", _pictUrl)
    return nil
}

// PictUrl Getter
func (r TaobaoSellercatsListAddRequest) GetPictUrl() string {
    return r._pictUrl
}
// ParentCid Setter
// 父类目编号，如果类目为店铺下的一级类目：值等于0，如果类目为子类目，调用获取taobao.sellercats.list.get父类目编号
func (r *TaobaoSellercatsListAddRequest) SetParentCid(_parentCid int64) error {
    r._parentCid = _parentCid
    r.Set("parent_cid", _parentCid)
    return nil
}

// ParentCid Getter
func (r TaobaoSellercatsListAddRequest) GetParentCid() int64 {
    return r._parentCid
}
// SortOrder Setter
// 该类目在页面上的排序位置,取值范围:大于零的整数
func (r *TaobaoSellercatsListAddRequest) SetSortOrder(_sortOrder int64) error {
    r._sortOrder = _sortOrder
    r.Set("sort_order", _sortOrder)
    return nil
}

// SortOrder Getter
func (r TaobaoSellercatsListAddRequest) GetSortOrder() int64 {
    return r._sortOrder
}
