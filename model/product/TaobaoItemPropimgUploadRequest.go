package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加或修改属性图片 API请求
taobao.item.propimg.upload

添加一张商品属性图片到num_iid指定的商品中 <br/>传入的num_iid所对应的商品必须属于当前会话的用户 <br/>图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 <br/>商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 <br/>商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。
*/
type TaobaoItemPropimgUploadRequest struct {
    model.Params
    // 商品数字ID，必选
    numIid   int64
    // 属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。
    properties   string
    // 属性图片内容。类型:JPG,GIF;图片大小不超过:3M
    image   []*model.File
    // 属性图片ID。如果是新增不需要填写
    id   int64
    // 图片位置
    position   int64
}

// 初始化TaobaoItemPropimgUploadRequest对象
func NewTaobaoItemPropimgUploadRequest() *TaobaoItemPropimgUploadRequest{
    return &TaobaoItemPropimgUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemPropimgUploadRequest) GetApiMethodName() string {
    return "taobao.item.propimg.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemPropimgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIid Setter
// 商品数字ID，必选
func (r *TaobaoItemPropimgUploadRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemPropimgUploadRequest) GetNumIid() int64 {
    return r.numIid
}
// Properties Setter
// 属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。
func (r *TaobaoItemPropimgUploadRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

// Properties Getter
func (r TaobaoItemPropimgUploadRequest) GetProperties() string {
    return r.properties
}
// Image Setter
// 属性图片内容。类型:JPG,GIF;图片大小不超过:3M
func (r *TaobaoItemPropimgUploadRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TaobaoItemPropimgUploadRequest) GetImage() []*model.File {
    return r.image
}
// Id Setter
// 属性图片ID。如果是新增不需要填写
func (r *TaobaoItemPropimgUploadRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoItemPropimgUploadRequest) GetId() int64 {
    return r.id
}
// Position Setter
// 图片位置
func (r *TaobaoItemPropimgUploadRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

// Position Getter
func (r TaobaoItemPropimgUploadRequest) GetPosition() int64 {
    return r.position
}
