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
    _img   []*model.File
    // 条形码
    _barCode   string
    // 商品ID，若填入商品ID则以商品ID为准，否则以outerId/barCode为准
    _itemId   int64
    // 商品outerId
    _outerId   string
    // 是否为主图
    _major   bool
    // 图片顺序
    _position   int64
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
func (r *TaobaoOmniitemItemImageUploadRequest) SetImg(_img []*model.File) error {
    r._img = _img
    r.Set("img", _img)
    return nil
}

// Img Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetImg() []*model.File {
    return r._img
}
// BarCode Setter
// 条形码
func (r *TaobaoOmniitemItemImageUploadRequest) SetBarCode(_barCode string) error {
    r._barCode = _barCode
    r.Set("bar_code", _barCode)
    return nil
}

// BarCode Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetBarCode() string {
    return r._barCode
}
// ItemId Setter
// 商品ID，若填入商品ID则以商品ID为准，否则以outerId/barCode为准
func (r *TaobaoOmniitemItemImageUploadRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetItemId() int64 {
    return r._itemId
}
// OuterId Setter
// 商品outerId
func (r *TaobaoOmniitemItemImageUploadRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetOuterId() string {
    return r._outerId
}
// Major Setter
// 是否为主图
func (r *TaobaoOmniitemItemImageUploadRequest) SetMajor(_major bool) error {
    r._major = _major
    r.Set("major", _major)
    return nil
}

// Major Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetMajor() bool {
    return r._major
}
// Position Setter
// 图片顺序
func (r *TaobaoOmniitemItemImageUploadRequest) SetPosition(_position int64) error {
    r._position = _position
    r.Set("position", _position)
    return nil
}

// Position Getter
func (r TaobaoOmniitemItemImageUploadRequest) GetPosition() int64 {
    return r._position
}
