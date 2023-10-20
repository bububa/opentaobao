package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoproductimguploadAPIRequest 上传单张产品非主图，如果需要传多张，可调多次 API请求
// taobao.product.img.upload
//
// 1.传入产品ID &lt;br/&gt;2.传入图片内容 &lt;br/&gt;注意：图片最大为500K,只支持JPG,GIF格式,如果需要传多张，可调多次
type TaobaoproductimguploadAPIRequest struct {
	model.Params
	// 产品图片ID.修改图片时需要传入
	_id int64
	// 产品ID.Product的id
	_productId int64
	// 图片内容.图片最大为500K,只支持JPG,GIF格式.
	_image *model.File
	// 图片序号
	_position int64
	// 是否将该图片设为主图.可选值:true,false;默认值:false.
	_isMajor bool
}

// NewTaobaoproductimguploadRequest 初始化TaobaoproductimguploadAPIRequest对象
func NewTaobaoproductimguploadRequest() *TaobaoproductimguploadAPIRequest {
	return &TaobaoproductimguploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoproductimguploadAPIRequest) GetApiMethodName() string {
	return "taobao.product.img.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoproductimguploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoproductimguploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 产品图片ID.修改图片时需要传入
func (r *TaobaoproductimguploadAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoproductimguploadAPIRequest) GetId() int64 {
	return r._id
}

// SetProductId is ProductId Setter
// 产品ID.Product的id
func (r *TaobaoproductimguploadAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoproductimguploadAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetImage is Image Setter
// 图片内容.图片最大为500K,只支持JPG,GIF格式.
func (r *TaobaoproductimguploadAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoproductimguploadAPIRequest) GetImage() *model.File {
	return r._image
}

// SetPosition is Position Setter
// 图片序号
func (r *TaobaoproductimguploadAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// GetPosition Position Getter
func (r TaobaoproductimguploadAPIRequest) GetPosition() int64 {
	return r._position
}

// SetIsMajor is IsMajor Setter
// 是否将该图片设为主图.可选值:true,false;默认值:false.
func (r *TaobaoproductimguploadAPIRequest) SetIsMajor(_isMajor bool) error {
	r._isMajor = _isMajor
	r.Set("is_major", _isMajor)
	return nil
}

// GetIsMajor IsMajor Getter
func (r TaobaoproductimguploadAPIRequest) GetIsMajor() bool {
	return r._isMajor
}
