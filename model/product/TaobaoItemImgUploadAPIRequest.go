package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemImgUploadAPIRequest 添加商品图片 API请求
// taobao.item.img.upload
//
// 添加一张商品图片到num_iid指定的商品中
// 传入的num_iid所对应的商品必须属于当前会话的用户
// 如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置 。
// 使用taobao.item.seller.get中返回的item_imgs字段获取图片id。
// 商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。
type TaobaoItemImgUploadAPIRequest struct {
	model.Params
	// 商品图片id(如果是更新图片，则需要传该参数)
	_id int64
	// 商品数字ID，该参数必须
	_numIid int64
	// 图片序号
	_position int64
	// 商品图片内容类型:JPG;最大:3M 。支持的文件类型：jpg,jpeg,png
	_image *model.File
	// 是否将该图片设为主图,可选值:true,false;默认值:false(非主图)
	_isMajor bool
	// 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
	_isRectangle bool
}

// NewTaobaoItemImgUploadRequest 初始化TaobaoItemImgUploadAPIRequest对象
func NewTaobaoItemImgUploadRequest() *TaobaoItemImgUploadAPIRequest {
	return &TaobaoItemImgUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemImgUploadAPIRequest) GetApiMethodName() string {
	return "taobao.item.img.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemImgUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 商品图片id(如果是更新图片，则需要传该参数)
func (r *TaobaoItemImgUploadAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r TaobaoItemImgUploadAPIRequest) GetId() int64 {
	return r._id
}

// Set is NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemImgUploadAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// Get NumIid Getter
func (r TaobaoItemImgUploadAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// Set is Position Setter
// 图片序号
func (r *TaobaoItemImgUploadAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// Get Position Getter
func (r TaobaoItemImgUploadAPIRequest) GetPosition() int64 {
	return r._position
}

// Set is Image Setter
// 商品图片内容类型:JPG;最大:3M 。支持的文件类型：jpg,jpeg,png
func (r *TaobaoItemImgUploadAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// Get Image Getter
func (r TaobaoItemImgUploadAPIRequest) GetImage() *model.File {
	return r._image
}

// Set is IsMajor Setter
// 是否将该图片设为主图,可选值:true,false;默认值:false(非主图)
func (r *TaobaoItemImgUploadAPIRequest) SetIsMajor(_isMajor bool) error {
	r._isMajor = _isMajor
	r.Set("is_major", _isMajor)
	return nil
}

// Get IsMajor Getter
func (r TaobaoItemImgUploadAPIRequest) GetIsMajor() bool {
	return r._isMajor
}

// Set is IsRectangle Setter
// 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
func (r *TaobaoItemImgUploadAPIRequest) SetIsRectangle(_isRectangle bool) error {
	r._isRectangle = _isRectangle
	r.Set("is_rectangle", _isRectangle)
	return nil
}

// Get IsRectangle Getter
func (r TaobaoItemImgUploadAPIRequest) GetIsRectangle() bool {
	return r._isRectangle
}
