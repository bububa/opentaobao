package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品上传图片 API请求
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

// 初始化TaobaoOmniitemItemImageUploadRequest对象
func NewTaobaoOmniitemItemImageUploadRequest() *TaobaoOmniitemItemImageUploadRequest{
    return &TaobaoOmniitemItemImageUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemImageUploadRequest) GetApiMethodName() string {
    return "taobao.omniitem.item.image.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Img Setter
// 商品图片信息，允许png、jpg、gif图片格式,3M以内
func (r *TaobaoOmniitemItemImageUploadRequest) SetImg(img []*model.File) error {
    r.img = img
    r.Set("img", img)
    return nil
}

// Img Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetImg() []*model.File {
    return r.img
}
// BarCode Setter
// 条形码
func (r *TaobaoOmniitemItemImageUploadRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

// BarCode Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetBarCode() string {
    return r.barCode
}
// ItemId Setter
// 商品ID，若填入商品ID则以商品ID为准，否则以outerId/barCode为准
func (r *TaobaoOmniitemItemImageUploadRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetItemId() int64 {
    return r.itemId
}
// OuterId Setter
// 商品outerId
func (r *TaobaoOmniitemItemImageUploadRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetOuterId() string {
    return r.outerId
}
// Major Setter
// 是否为主图
func (r *TaobaoOmniitemItemImageUploadRequest) SetMajor(major bool) error {
    r.major = major
    r.Set("major", major)
    return nil
}

// Major Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetMajor() bool {
    return r.major
}
// Position Setter
// 图片顺序
func (r *TaobaoOmniitemItemImageUploadRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

// Position Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetPosition() int64 {
    return r.position
}
