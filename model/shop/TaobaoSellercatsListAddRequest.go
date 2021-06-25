package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加卖家自定义类目 APIRequest
taobao.sellercats.list.add

此API添加卖家店铺内自定义类目 <br/>父类目parent_cid值等于0：表示此类目为店铺下的一级类目，值不等于0：表示此类目有父类目 <br/>注：因为缓存的关系,添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
*/
type TaobaoSellercatsListAddRequest struct {
    model.Params

    // 卖家自定义类目名称。不超过20个字符
    name   string 

    // 链接图片URL地址。(绝对地址，格式：http://host/image_path)
    pictUrl   string 

    // 父类目编号，如果类目为店铺下的一级类目：值等于0，如果类目为子类目，调用获取taobao.sellercats.list.get父类目编号
    parentCid   int64 

    // 该类目在页面上的排序位置,取值范围:大于零的整数
    sortOrder   int64 

}

func NewTaobaoSellercatsListAddRequest() *TaobaoSellercatsListAddRequest{
    return &TaobaoSellercatsListAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSellercatsListAddRequest) GetApiMethodName() string {
    return "taobao.sellercats.list.add"
}

func (r TaobaoSellercatsListAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSellercatsListAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoSellercatsListAddRequest) GetName() string {
    return r.name
}

func (r *TaobaoSellercatsListAddRequest) SetPictUrl(pictUrl string) error {
    r.pictUrl = pictUrl
    r.Set("pict_url", pictUrl)
    return nil
}

func (r TaobaoSellercatsListAddRequest) GetPictUrl() string {
    return r.pictUrl
}

func (r *TaobaoSellercatsListAddRequest) SetParentCid(parentCid int64) error {
    r.parentCid = parentCid
    r.Set("parent_cid", parentCid)
    return nil
}

func (r TaobaoSellercatsListAddRequest) GetParentCid() int64 {
    return r.parentCid
}

func (r *TaobaoSellercatsListAddRequest) SetSortOrder(sortOrder int64) error {
    r.sortOrder = sortOrder
    r.Set("sort_order", sortOrder)
    return nil
}

func (r TaobaoSellercatsListAddRequest) GetSortOrder() int64 {
    return r.sortOrder
}

