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
    productId   int64
    // 产品主图图片空间相对路径或绝对路径
    picPath   string
    // 产品图片
    image   []*model.File
    // 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
    position   int64
    // properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
    properties   string
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
func (r *TaobaoFenxiaoProductImageUploadRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductImageUploadRequest) GetProductId() int64 {
    return r.productId
}
// PicPath Setter
// 产品主图图片空间相对路径或绝对路径
func (r *TaobaoFenxiaoProductImageUploadRequest) SetPicPath(picPath string) error {
    r.picPath = picPath
    r.Set("pic_path", picPath)
    return nil
}

// PicPath Getter
func (r TaobaoFenxiaoProductImageUploadRequest) GetPicPath() string {
    return r.picPath
}
// Image Setter
// 产品图片
func (r *TaobaoFenxiaoProductImageUploadRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TaobaoFenxiaoProductImageUploadRequest) GetImage() []*model.File {
    return r.image
}
// Position Setter
// 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
func (r *TaobaoFenxiaoProductImageUploadRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

// Position Getter
func (r TaobaoFenxiaoProductImageUploadRequest) GetPosition() int64 {
    return r.position
}
// Properties Setter
// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
func (r *TaobaoFenxiaoProductImageUploadRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

// Properties Getter
func (r TaobaoFenxiaoProductImageUploadRequest) GetProperties() string {
    return r.properties
}
