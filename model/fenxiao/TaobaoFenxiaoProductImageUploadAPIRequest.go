package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductimageuploadAPIRequest 产品图片上传 API请求
// taobao.fenxiao.product.image.upload
//
// 产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片
type TaobaofenxiaoproductimageuploadAPIRequest struct {
	model.Params
	// 产品主图图片空间相对路径或绝对路径
	_picPath string
	// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
	_properties string
	// 产品ID
	_productId int64
	// 产品图片
	_image *model.File
	// 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
	_position int64
}

// NewTaobaofenxiaoproductimageuploadRequest 初始化TaobaofenxiaoproductimageuploadAPIRequest对象
func NewTaobaofenxiaoproductimageuploadRequest() *TaobaofenxiaoproductimageuploadAPIRequest {
	return &TaobaofenxiaoproductimageuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductimageuploadAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductimageuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductimageuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPicPath is PicPath Setter
// 产品主图图片空间相对路径或绝对路径
func (r *TaobaofenxiaoproductimageuploadAPIRequest) SetPicPath(_picPath string) error {
	r._picPath = _picPath
	r.Set("pic_path", _picPath)
	return nil
}

// GetPicPath PicPath Getter
func (r TaobaofenxiaoproductimageuploadAPIRequest) GetPicPath() string {
	return r._picPath
}

// SetProperties is Properties Setter
// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
func (r *TaobaofenxiaoproductimageuploadAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaofenxiaoproductimageuploadAPIRequest) GetProperties() string {
	return r._properties
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaofenxiaoproductimageuploadAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaofenxiaoproductimageuploadAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetImage is Image Setter
// 产品图片
func (r *TaobaofenxiaoproductimageuploadAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaofenxiaoproductimageuploadAPIRequest) GetImage() *model.File {
	return r._image
}

// SetPosition is Position Setter
// 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
func (r *TaobaofenxiaoproductimageuploadAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// GetPosition Position Getter
func (r TaobaofenxiaoproductimageuploadAPIRequest) GetPosition() int64 {
	return r._position
}
