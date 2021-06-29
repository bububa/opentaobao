package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品上传图片 APIRequest
taobao.omniitem.item.image.upload

全渠道商品上传图片
*/
type TaobaoOmniitemItemImageUploadRequest struct {
    model.Params

    // 商品图片信息，允许png、jpg、gif图片格式,3M以内
    img   []*model.File 

    // 条形码
    barCode   string 

    // 商品ID，若填入商品ID则以商品ID为准，否则以outerId/barCode为准
    itemId   int64 

    // 商品outerId
    outerId   string 

    // 是否为主图
    major   bool 

    // 图片顺序
    position   int64 

}

func NewTaobaoOmniitemItemImageUploadRequest() *TaobaoOmniitemItemImageUploadRequest{
    return &TaobaoOmniitemItemImageUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemItemImageUploadRequest) GetApiMethodName() string {
    return "taobao.omniitem.item.image.upload"
}

func (r TaobaoOmniitemItemImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemItemImageUploadRequest) SetImg(img []*model.File) error {
    r.img = img
    r.Set("img", img)
    return nil
}

func (r TaobaoOmniitemItemImageUploadRequest) GetImg() []*model.File {
    return r.img
}

func (r *TaobaoOmniitemItemImageUploadRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

func (r TaobaoOmniitemItemImageUploadRequest) GetBarCode() string {
    return r.barCode
}

func (r *TaobaoOmniitemItemImageUploadRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOmniitemItemImageUploadRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoOmniitemItemImageUploadRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoOmniitemItemImageUploadRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoOmniitemItemImageUploadRequest) SetMajor(major bool) error {
    r.major = major
    r.Set("major", major)
    return nil
}

func (r TaobaoOmniitemItemImageUploadRequest) GetMajor() bool {
    return r.major
}

func (r *TaobaoOmniitemItemImageUploadRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

func (r TaobaoOmniitemItemImageUploadRequest) GetPosition() int64 {
    return r.position
}

