package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单张产品非主图，如果需要传多张，可调多次 API请求
taobao.product.img.upload

1.传入产品ID <br/>2.传入图片内容 <br/>注意：图片最大为500K,只支持JPG,GIF格式,如果需要传多张，可调多次
*/
type TaobaoProductImgUploadRequest struct {
    model.Params
    // 产品图片ID.修改图片时需要传入
    _id   int64
    // 产品ID.Product的id
    _productId   int64
    // 图片内容.图片最大为500K,只支持JPG,GIF格式.
    _image   *model.File
    // 图片序号
    _position   int64
    // 是否将该图片设为主图.可选值:true,false;默认值:false.
    _isMajor   bool
}

// 初始化TaobaoProductImgUploadRequest对象
func NewTaobaoProductImgUploadRequest() *TaobaoProductImgUploadRequest{
    return &TaobaoProductImgUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoProductImgUploadRequest) GetApiMethodName() string {
    return "taobao.product.img.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoProductImgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 产品图片ID.修改图片时需要传入
func (r *TaobaoProductImgUploadRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TaobaoProductImgUploadRequest) GetId() int64 {
    return r._id
}
// ProductId Setter
// 产品ID.Product的id
func (r *TaobaoProductImgUploadRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoProductImgUploadRequest) GetProductId() int64 {
    return r._productId
}
// Image Setter
// 图片内容.图片最大为500K,只支持JPG,GIF格式.
func (r *TaobaoProductImgUploadRequest) SetImage(_image *model.File) error {
    r._image = _image
    r.Set("image", _image)
    return nil
}

// Image Getter
func (r TaobaoProductImgUploadRequest) GetImage() *model.File {
    return r._image
}
// Position Setter
// 图片序号
func (r *TaobaoProductImgUploadRequest) SetPosition(_position int64) error {
    r._position = _position
    r.Set("position", _position)
    return nil
}

// Position Getter
func (r TaobaoProductImgUploadRequest) GetPosition() int64 {
    return r._position
}
// IsMajor Setter
// 是否将该图片设为主图.可选值:true,false;默认值:false.
func (r *TaobaoProductImgUploadRequest) SetIsMajor(_isMajor bool) error {
    r._isMajor = _isMajor
    r.Set("is_major", _isMajor)
    return nil
}

// IsMajor Getter
func (r TaobaoProductImgUploadRequest) GetIsMajor() bool {
    return r._isMajor
}
