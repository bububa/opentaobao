package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建IC商品与后端商品的映射关系 APIRequest
taobao.scitem.map.add

创建IC商品或分销商品与后端商品的映射关系
*/
type TaobaoScitemMapAddRequest struct {
    model.Params

    // 前台ic商品id
    itemId   int64 

    // 前台ic商品skuid
    skuId   int64 

    // sc_item_id和outer_code 其中一个不能为空
    scItemId   int64 

    // sc_item_id和outer_code 其中一个不能为空
    outerCode   string 

    // 默认值为false<br/>true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联<br/>false:不进行高级校验
    needCheck   bool 

}

func NewTaobaoScitemMapAddRequest() *TaobaoScitemMapAddRequest{
    return &TaobaoScitemMapAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoScitemMapAddRequest) GetApiMethodName() string {
    return "taobao.scitem.map.add"
}

func (r TaobaoScitemMapAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoScitemMapAddRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoScitemMapAddRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoScitemMapAddRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoScitemMapAddRequest) GetSkuId() int64 {
    return r.skuId
}

func (r *TaobaoScitemMapAddRequest) SetScItemId(scItemId int64) error {
    r.scItemId = scItemId
    r.Set("sc_item_id", scItemId)
    return nil
}

func (r TaobaoScitemMapAddRequest) GetScItemId() int64 {
    return r.scItemId
}

func (r *TaobaoScitemMapAddRequest) SetOuterCode(outerCode string) error {
    r.outerCode = outerCode
    r.Set("outer_code", outerCode)
    return nil
}

func (r TaobaoScitemMapAddRequest) GetOuterCode() string {
    return r.outerCode
}

func (r *TaobaoScitemMapAddRequest) SetNeedCheck(needCheck bool) error {
    r.needCheck = needCheck
    r.Set("need_check", needCheck)
    return nil
}

func (r TaobaoScitemMapAddRequest) GetNeedCheck() bool {
    return r.needCheck
}

