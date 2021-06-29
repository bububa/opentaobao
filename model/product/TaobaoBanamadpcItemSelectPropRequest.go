package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取子属性 API请求
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

// 初始化TaobaoBanamadpcItemSelectPropRequest对象
func NewTaobaoBanamadpcItemSelectPropRequest() *TaobaoBanamadpcItemSelectPropRequest{
    return &TaobaoBanamadpcItemSelectPropRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemSelectPropRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.select.prop"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemSelectPropRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Xml Setter
// 子属性的schema xml
func (r *TaobaoBanamadpcItemSelectPropRequest) SetXml(xml string) error {
    r.xml = xml
    r.Set("xml", xml)
    return nil
}

// Xml Getter
func (r TaobaoBanamadpcItemSelectPropRequest) GetXml() string {
    return r.xml
}
// PropId Setter
// 属性id
func (r *TaobaoBanamadpcItemSelectPropRequest) SetPropId(propId int64) error {
    r.propId = propId
    r.Set("prop_id", propId)
    return nil
}

// PropId Getter
func (r TaobaoBanamadpcItemSelectPropRequest) GetPropId() int64 {
    return r.propId
}
// CatId Setter
// 类目id
func (r *TaobaoBanamadpcItemSelectPropRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r TaobaoBanamadpcItemSelectPropRequest) GetCatId() int64 {
    return r.catId
}
