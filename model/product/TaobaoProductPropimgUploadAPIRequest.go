package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductPropimgUploadAPIRequest 上传单张产品属性图片，如果需要传多张，可调多次 API请求
// taobao.product.propimg.upload
//
// 传入产品ID &lt;br/&gt;传入props,目前仅支持颜色属性.调用taobao.itemprops.get.v2取得颜色属性pid,&lt;br/&gt;再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串; &lt;br/&gt;传入图片内容 &lt;br/&gt;注意：图片最大为2M,只支持JPG,GIF,如果需要传多张，可调多次
type TaobaoProductPropimgUploadAPIRequest struct {
	model.Params
	// 属性串.目前仅支持颜色属性.调用taobao.itemprops.get获取类目属性,取得颜色属性pid,再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串;
	_props string
	// 产品属性图片ID
	_id int64
	// 产品ID.Product的id
	_productId int64
	// 图片内容.图片最大为2M,只支持JPG,GIF.
	_image *model.File
	// 图片序号
	_position int64
}

// NewTaobaoProductPropimgUploadRequest 初始化TaobaoProductPropimgUploadAPIRequest对象
func NewTaobaoProductPropimgUploadRequest() *TaobaoProductPropimgUploadAPIRequest {
	return &TaobaoProductPropimgUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoProductPropimgUploadAPIRequest) GetApiMethodName() string {
	return "taobao.product.propimg.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoProductPropimgUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoProductPropimgUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProps is Props Setter
// 属性串.目前仅支持颜色属性.调用taobao.itemprops.get获取类目属性,取得颜色属性pid,再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串;
func (r *TaobaoProductPropimgUploadAPIRequest) SetProps(_props string) error {
	r._props = _props
	r.Set("props", _props)
	return nil
}

// GetProps Props Getter
func (r TaobaoProductPropimgUploadAPIRequest) GetProps() string {
	return r._props
}

// SetId is Id Setter
// 产品属性图片ID
func (r *TaobaoProductPropimgUploadAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoProductPropimgUploadAPIRequest) GetId() int64 {
	return r._id
}

// SetProductId is ProductId Setter
// 产品ID.Product的id
func (r *TaobaoProductPropimgUploadAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoProductPropimgUploadAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetImage is Image Setter
// 图片内容.图片最大为2M,只支持JPG,GIF.
func (r *TaobaoProductPropimgUploadAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoProductPropimgUploadAPIRequest) GetImage() *model.File {
	return r._image
}

// SetPosition is Position Setter
// 图片序号
func (r *TaobaoProductPropimgUploadAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// GetPosition Position Getter
func (r TaobaoProductPropimgUploadAPIRequest) GetPosition() int64 {
	return r._position
}
