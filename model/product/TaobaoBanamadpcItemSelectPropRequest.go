package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取子属性 APIRequest
taobao.banamadpc.item.select.prop

巴拿马供应商通过此接口获取子属性
*/
type TaobaoBanamadpcItemSelectPropRequest struct {
    model.Params

    // 子属性的schema xml
    xml   string 

    // 属性id
    propId   int64 

    // 类目id
    catId   int64 

}

func NewTaobaoBanamadpcItemSelectPropRequest() *TaobaoBanamadpcItemSelectPropRequest{
    return &TaobaoBanamadpcItemSelectPropRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBanamadpcItemSelectPropRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.select.prop"
}

func (r TaobaoBanamadpcItemSelectPropRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBanamadpcItemSelectPropRequest) SetXml(xml string) error {
    r.xml = xml
    r.Set("xml", xml)
    return nil
}

func (r TaobaoBanamadpcItemSelectPropRequest) GetXml() string {
    return r.xml
}

func (r *TaobaoBanamadpcItemSelectPropRequest) SetPropId(propId int64) error {
    r.propId = propId
    r.Set("prop_id", propId)
    return nil
}

func (r TaobaoBanamadpcItemSelectPropRequest) GetPropId() int64 {
    return r.propId
}

func (r *TaobaoBanamadpcItemSelectPropRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r TaobaoBanamadpcItemSelectPropRequest) GetCatId() int64 {
    return r.catId
}

