package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品图片上传 APIRequest
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
    image   []byte 

    // 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
    position   int64 

    // properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
    properties   string 

}

func NewTaobaoFenxiaoProductImageUploadRequest() *TaobaoFenxiaoProductImageUploadRequest{
    return &TaobaoFenxiaoProductImageUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductImageUploadRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.image.upload"
}

func (r TaobaoFenxiaoProductImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductImageUploadRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoFenxiaoProductImageUploadRequest) GetProductId() int64 {
    return r.productId
}

func (r *TaobaoFenxiaoProductImageUploadRequest) SetPicPath(picPath string) error {
    r.picPath = picPath
    r.Set("pic_path", picPath)
    return nil
}

func (r TaobaoFenxiaoProductImageUploadRequest) GetPicPath() string {
    return r.picPath
}

func (r *TaobaoFenxiaoProductImageUploadRequest) SetImage(image []byte) error {
    r.image = image
    r.Set("image", image)
    return nil
}

func (r TaobaoFenxiaoProductImageUploadRequest) GetImage() []byte {
    return r.image
}

func (r *TaobaoFenxiaoProductImageUploadRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

func (r TaobaoFenxiaoProductImageUploadRequest) GetPosition() int64 {
    return r.position
}

func (r *TaobaoFenxiaoProductImageUploadRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

func (r TaobaoFenxiaoProductImageUploadRequest) GetProperties() string {
    return r.properties
}

