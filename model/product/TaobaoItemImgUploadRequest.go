package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加商品图片 API请求
taobao.item.img.upload

添加一张商品图片到num_iid指定的商品中 
传入的num_iid所对应的商品必须属于当前会话的用户 
如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置 。
使用taobao.item.seller.get中返回的item_imgs字段获取图片id。
商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。
*/
type TaobaoItemImgUploadRequest struct {
    model.Params
    // 商品图片id(如果是更新图片，则需要传该参数)
    id   int64
    // 商品数字ID，该参数必须
    numIid   int64
    // 图片序号
    position   int64
    // 商品图片内容类型:JPG;最大:3M 。支持的文件类型：jpg,jpeg,png
    image   []*model.File
    // 是否将该图片设为主图,可选值:true,false;默认值:false(非主图)
    isMajor   bool
    // 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
    isRectangle   bool
}

// 初始化TaobaoItemImgUploadRequest对象
func NewTaobaoItemImgUploadRequest() *TaobaoItemImgUploadRequest{
    return &TaobaoItemImgUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemImgUploadRequest) GetApiMethodName() string {
    return "taobao.item.img.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemImgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 商品图片id(如果是更新图片，则需要传该参数)
func (r *TaobaoItemImgUploadRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoItemImgUploadRequest) GetId() int64 {
    return r.id
}
// NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemImgUploadRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemImgUploadRequest) GetNumIid() int64 {
    return r.numIid
}
// Position Setter
// 图片序号
func (r *TaobaoItemImgUploadRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

// Position Getter
func (r TaobaoItemImgUploadRequest) GetPosition() int64 {
    return r.position
}
// Image Setter
// 商品图片内容类型:JPG;最大:3M 。支持的文件类型：jpg,jpeg,png
func (r *TaobaoItemImgUploadRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TaobaoItemImgUploadRequest) GetImage() []*model.File {
    return r.image
}
// IsMajor Setter
// 是否将该图片设为主图,可选值:true,false;默认值:false(非主图)
func (r *TaobaoItemImgUploadRequest) SetIsMajor(isMajor bool) error {
    r.isMajor = isMajor
    r.Set("is_major", isMajor)
    return nil
}

// IsMajor Getter
func (r TaobaoItemImgUploadRequest) GetIsMajor() bool {
    return r.isMajor
}
// IsRectangle Setter
// 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
func (r *TaobaoItemImgUploadRequest) SetIsRectangle(isRectangle bool) error {
    r.isRectangle = isRectangle
    r.Set("is_rectangle", isRectangle)
    return nil
}

// IsRectangle Getter
func (r TaobaoItemImgUploadRequest) GetIsRectangle() bool {
    return r.isRectangle
}
