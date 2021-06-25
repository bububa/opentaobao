package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取可用宝贝描述规范化模块 APIRequest
taobao.item.anchor.get

根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点
*/
type TaobaoItemAnchorGetRequest struct {
    model.Params

    // 宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1.(最小值：-1，最大值：1)
    type   int64 

    // 对应类目编号
    catId   int64 

}

func NewTaobaoItemAnchorGetRequest() *TaobaoItemAnchorGetRequest{
    return &TaobaoItemAnchorGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemAnchorGetRequest) GetApiMethodName() string {
    return "taobao.item.anchor.get"
}

func (r TaobaoItemAnchorGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemAnchorGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoItemAnchorGetRequest) GetType() int64 {
    return r.type
}

func (r *TaobaoItemAnchorGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r TaobaoItemAnchorGetRequest) GetCatId() int64 {
    return r.catId
}

