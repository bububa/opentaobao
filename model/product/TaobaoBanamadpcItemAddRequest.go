package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新发商品 API请求
taobao.banamadpc.item.add

巴拿马供应商通过此接口新发商品
*/
type TaobaoBanamadpcItemAddRequest struct {
    model.Params
    // 类目id
    catId   int64
    // 商品的schema xml
    xml   string
}

// 初始化TaobaoBanamadpcItemAddRequest对象
func NewTaobaoBanamadpcItemAddRequest() *TaobaoBanamadpcItemAddRequest{
    return &TaobaoBanamadpcItemAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemAddRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 类目id
func (r *TaobaoBanamadpcItemAddRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r TaobaoBanamadpcItemAddRequest) GetCatId() int64 {
    return r.catId
}
// Xml Setter
// 商品的schema xml
func (r *TaobaoBanamadpcItemAddRequest) SetXml(xml string) error {
    r.xml = xml
    r.Set("xml", xml)
    return nil
}

// Xml Getter
func (r TaobaoBanamadpcItemAddRequest) GetXml() string {
    return r.xml
}
