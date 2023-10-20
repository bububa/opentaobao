package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductImgUploadAPIRequest 上传单张产品非主图，如果需要传多张，可调多次 API请求
// taobao.product.img.upload
//
// 1.传入产品ID &lt;br/&gt;2.传入图片内容 &lt;br/&gt;注意：图片最大为500K,只支持JPG,GIF格式,如果需要传多张，可调多次
type TaobaoProductImgUploadAPIRequest struct {
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

// NewTaobaoProductImgUploadRequest 初始化TaobaoProductImgUploadAPIRequest对象
func NewTaobaoProductImgUploadRequest() *TaobaoProductImgUploadAPIRequest {
	return &TaobaoProductImgUploadAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoProductImgUploadAPIRequest) Reset() {
	r._id = 0
	r._productId = 0
	r._image = nil
	r._position = 0
	r._isMajor = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoProductImgUploadAPIRequest) GetApiMethodName() string {
	return "taobao.product.img.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoProductImgUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoProductImgUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 产品图片ID.修改图片时需要传入
func (r *TaobaoProductImgUploadAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoProductImgUploadAPIRequest) GetId() int64 {
	return r._id
}

// SetProductId is ProductId Setter
// 产品ID.Product的id
func (r *TaobaoProductImgUploadAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoProductImgUploadAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetImage is Image Setter
// 图片内容.图片最大为500K,只支持JPG,GIF格式.
func (r *TaobaoProductImgUploadAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoProductImgUploadAPIRequest) GetImage() *model.File {
	return r._image
}

// SetPosition is Position Setter
// 图片序号
func (r *TaobaoProductImgUploadAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// GetPosition Position Getter
func (r TaobaoProductImgUploadAPIRequest) GetPosition() int64 {
	return r._position
}

// SetIsMajor is IsMajor Setter
// 是否将该图片设为主图.可选值:true,false;默认值:false.
func (r *TaobaoProductImgUploadAPIRequest) SetIsMajor(_isMajor bool) error {
	r._isMajor = _isMajor
	r.Set("is_major", _isMajor)
	return nil
}

// GetIsMajor IsMajor Getter
func (r TaobaoProductImgUploadAPIRequest) GetIsMajor() bool {
	return r._isMajor
}

var poolTaobaoProductImgUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoProductImgUploadRequest()
	},
}

// GetTaobaoProductImgUploadRequest 从 sync.Pool 获取 TaobaoProductImgUploadAPIRequest
func GetTaobaoProductImgUploadAPIRequest() *TaobaoProductImgUploadAPIRequest {
	return poolTaobaoProductImgUploadAPIRequest.Get().(*TaobaoProductImgUploadAPIRequest)
}

// ReleaseTaobaoProductImgUploadAPIRequest 将 TaobaoProductImgUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoProductImgUploadAPIRequest(v *TaobaoProductImgUploadAPIRequest) {
	v.Reset()
	poolTaobaoProductImgUploadAPIRequest.Put(v)
}
