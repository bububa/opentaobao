package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新发商品 APIRequest
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

func NewTaobaoBanamadpcItemAddRequest() *TaobaoBanamadpcItemAddRequest{
    return &TaobaoBanamadpcItemAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBanamadpcItemAddRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.add"
}

func (r TaobaoBanamadpcItemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBanamadpcItemAddRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r TaobaoBanamadpcItemAddRequest) GetCatId() int64 {
    return r.catId
}

func (r *TaobaoBanamadpcItemAddRequest) SetXml(xml string) error {
    r.xml = xml
    r.Set("xml", xml)
    return nil
}

func (r TaobaoBanamadpcItemAddRequest) GetXml() string {
    return r.xml
}

