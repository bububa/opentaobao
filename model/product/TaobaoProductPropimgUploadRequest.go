package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单张产品属性图片，如果需要传多张，可调多次 API请求
taobao.product.propimg.upload

传入产品ID <br/>传入props,目前仅支持颜色属性.调用taobao.itemprops.get.v2取得颜色属性pid,<br/>再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串; <br/>传入图片内容 <br/>注意：图片最大为2M,只支持JPG,GIF,如果需要传多张，可调多次
*/
type TaobaoProductPropimgUploadRequest struct {
    model.Params
    // 产品属性图片ID
    _id   int64
    // 产品ID.Product的id
    _productId   int64
    // 属性串.目前仅支持颜色属性.调用taobao.itemprops.get获取类目属性,取得颜色属性pid,再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串;
    _props   string
    // 图片内容.图片最大为2M,只支持JPG,GIF.
    _image   []*model.File
    // 图片序号
    _position   int64
}

// 初始化TaobaoProductPropimgUploadRequest对象
func NewTaobaoProductPropimgUploadRequest() *TaobaoProductPropimgUploadRequest{
    return &TaobaoProductPropimgUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoProductPropimgUploadRequest) GetApiMethodName() string {
    return "taobao.product.propimg.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoProductPropimgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 产品属性图片ID
func (r *TaobaoProductPropimgUploadRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TaobaoProductPropimgUploadRequest) GetId() int64 {
    return r._id
}
// ProductId Setter
// 产品ID.Product的id
func (r *TaobaoProductPropimgUploadRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoProductPropimgUploadRequest) GetProductId() int64 {
    return r._productId
}
// Props Setter
// 属性串.目前仅支持颜色属性.调用taobao.itemprops.get获取类目属性,取得颜色属性pid,再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串;
func (r *TaobaoProductPropimgUploadRequest) SetProps(_props string) error {
    r._props = _props
    r.Set("props", _props)
    return nil
}

// Props Getter
func (r TaobaoProductPropimgUploadRequest) GetProps() string {
    return r._props
}
// Image Setter
// 图片内容.图片最大为2M,只支持JPG,GIF.
func (r *TaobaoProductPropimgUploadRequest) SetImage(_image []*model.File) error {
    r._image = _image
    r.Set("image", _image)
    return nil
}

// Image Getter
func (r TaobaoProductPropimgUploadRequest) GetImage() []*model.File {
    return r._image
}
// Position Setter
// 图片序号
func (r *TaobaoProductPropimgUploadRequest) SetPosition(_position int64) error {
    r._position = _position
    r.Set("position", _position)
    return nil
}

// Position Getter
func (r TaobaoProductPropimgUploadRequest) GetPosition() int64 {
    return r._position
}
