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
    id   int64
    // 产品ID.Product的id
    productId   int64
    // 属性串.目前仅支持颜色属性.调用taobao.itemprops.get获取类目属性,取得颜色属性pid,再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串;
    props   string
    // 图片内容.图片最大为2M,只支持JPG,GIF.
    image   []*model.File
    // 图片序号
    position   int64
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
func (r *TaobaoProductPropimgUploadRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoProductPropimgUploadRequest) GetId() int64 {
    return r.id
}
// ProductId Setter
// 产品ID.Product的id
func (r *TaobaoProductPropimgUploadRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoProductPropimgUploadRequest) GetProductId() int64 {
    return r.productId
}
// Props Setter
// 属性串.目前仅支持颜色属性.调用taobao.itemprops.get获取类目属性,取得颜色属性pid,再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串;
func (r *TaobaoProductPropimgUploadRequest) SetProps(props string) error {
    r.props = props
    r.Set("props", props)
    return nil
}

// Props Getter
func (r TaobaoProductPropimgUploadRequest) GetProps() string {
    return r.props
}
// Image Setter
// 图片内容.图片最大为2M,只支持JPG,GIF.
func (r *TaobaoProductPropimgUploadRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TaobaoProductPropimgUploadRequest) GetImage() []*model.File {
    return r.image
}
// Position Setter
// 图片序号
func (r *TaobaoProductPropimgUploadRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

// Position Getter
func (r TaobaoProductPropimgUploadRequest) GetPosition() int64 {
    return r.position
}
