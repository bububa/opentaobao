package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品图片上传 API请求
taobao.fenxiao.product.image.upload

产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片
*/
type TaobaoFenxiaoProductImageUploadRequest struct {
    model.Params
    // 产品ID
    _productId   int64
    // 产品主图图片空间相对路径或绝对路径
    _picPath   string
    // 产品图片
    _image   []*model.File
    // 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
    _position   int64
    // properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
    _properties   string
}

// 初始化TaobaoFenxiaoProductImageUploadRequest对象
func NewTaobaoFenxiaoProductImageUploadRequest() *TaobaoFenxiaoProductImageUploadRequest{
    return &TaobaoFenxiaoProductImageUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductImageUploadRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.image.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductImageUploadRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductImageUploadRequest) GetProductId() int64 {
    return r._productId
}
// PicPath Setter
// 产品主图图片空间相对路径或绝对路径
func (r *TaobaoFenxiaoProductImageUploadRequest) SetPicPath(_picPath string) error {
    r._picPath = _picPath
    r.Set("pic_path", _picPath)
    return nil
}

// PicPath Getter
func (r TaobaoFenxiaoProductImageUploadRequest) GetPicPath() string {
    return r._picPath
}
// Image Setter
// 产品图片
func (r *TaobaoFenxiaoProductImageUploadRequest) SetImage(_image []*model.File) error {
    r._image = _image
    r.Set("image", _image)
    return nil
}

// Image Getter
func (r TaobaoFenxiaoProductImageUploadRequest) GetImage() []*model.File {
    return r._image
}
// Position Setter
// 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
func (r *TaobaoFenxiaoProductImageUploadRequest) SetPosition(_position int64) error {
    r._position = _position
    r.Set("position", _position)
    return nil
}

// Position Getter
func (r TaobaoFenxiaoProductImageUploadRequest) GetPosition() int64 {
    return r._position
}
// Properties Setter
// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
func (r *TaobaoFenxiaoProductImageUploadRequest) SetProperties(_properties string) error {
    r._properties = _properties
    r.Set("properties", _properties)
    return nil
}

// Properties Getter
func (r TaobaoFenxiaoProductImageUploadRequest) GetProperties() string {
    return r._properties
}
