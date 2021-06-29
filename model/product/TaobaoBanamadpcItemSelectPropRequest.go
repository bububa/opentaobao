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
    _xml   string
    // 属性id
    _propId   int64
    // 类目id
    _catId   int64
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
func (r *TaobaoBanamadpcItemSelectPropRequest) SetXml(_xml string) error {
    r._xml = _xml
    r.Set("xml", _xml)
    return nil
}

// Xml Getter
func (r TaobaoBanamadpcItemSelectPropRequest) GetXml() string {
    return r._xml
}
// PropId Setter
// 属性id
func (r *TaobaoBanamadpcItemSelectPropRequest) SetPropId(_propId int64) error {
    r._propId = _propId
    r.Set("prop_id", _propId)
    return nil
}

// PropId Getter
func (r TaobaoBanamadpcItemSelectPropRequest) GetPropId() int64 {
    return r._propId
}
// CatId Setter
// 类目id
func (r *TaobaoBanamadpcItemSelectPropRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r TaobaoBanamadpcItemSelectPropRequest) GetCatId() int64 {
    return r._catId
}
