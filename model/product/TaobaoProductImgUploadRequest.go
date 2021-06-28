package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单张产品非主图，如果需要传多张，可调多次 APIRequest
taobao.product.img.upload

1.传入产品ID <br/>2.传入图片内容 <br/>注意：图片最大为500K,只支持JPG,GIF格式,如果需要传多张，可调多次
*/
type TaobaoProductImgUploadRequest struct {
    model.Params

    // 产品图片ID.修改图片时需要传入
    id   int64 

    // 产品ID.Product的id
    productId   int64 

    // 图片内容.图片最大为500K,只支持JPG,GIF格式.
    image   []*model.File 

    // 图片序号
    position   int64 

    // 是否将该图片设为主图.可选值:true,false;默认值:false.
    isMajor   bool 

}

func NewTaobaoProductImgUploadRequest() *TaobaoProductImgUploadRequest{
    return &TaobaoProductImgUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoProductImgUploadRequest) GetApiMethodName() string {
    return "taobao.product.img.upload"
}

func (r TaobaoProductImgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoProductImgUploadRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TaobaoProductImgUploadRequest) GetId() int64 {
    return r.id
}

func (r *TaobaoProductImgUploadRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoProductImgUploadRequest) GetProductId() int64 {
    return r.productId
}

func (r *TaobaoProductImgUploadRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

func (r TaobaoProductImgUploadRequest) GetImage() []*model.File {
    return r.image
}

func (r *TaobaoProductImgUploadRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

func (r TaobaoProductImgUploadRequest) GetPosition() int64 {
    return r.position
}

func (r *TaobaoProductImgUploadRequest) SetIsMajor(isMajor bool) error {
    r.isMajor = isMajor
    r.Set("is_major", isMajor)
    return nil
}

func (r TaobaoProductImgUploadRequest) GetIsMajor() bool {
    return r.isMajor
}

