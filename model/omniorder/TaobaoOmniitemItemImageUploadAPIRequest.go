package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemItemImageUploadAPIRequest 全渠道商品上传图片 API请求
// taobao.omniitem.item.image.upload
//
// 全渠道商品上传图片
type TaobaoOmniitemItemImageUploadAPIRequest struct {
	model.Params
	// 商品图片信息，允许png、jpg、gif图片格式,3M以内
	_img *model.File
	// 条形码
	_barCode string
	// 商品ID，若填入商品ID则以商品ID为准，否则以outerId/barCode为准
	_itemId int64
	// 商品outerId
	_outerId string
	// 是否为主图
	_major bool
	// 图片顺序
	_position int64
}

// NewTaobaoOmniitemItemImageUploadRequest 初始化TaobaoOmniitemItemImageUploadAPIRequest对象
func NewTaobaoOmniitemItemImageUploadRequest() *TaobaoOmniitemItemImageUploadAPIRequest {
	return &TaobaoOmniitemItemImageUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemImageUploadAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.item.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemImageUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Img Setter
// 商品图片信息，允许png、jpg、gif图片格式,3M以内
func (r *TaobaoOmniitemItemImageUploadAPIRequest) SetImg(_img *model.File) error {
	r._img = _img
	r.Set("img", _img)
	return nil
}

// Get Img Getter
func (r TaobaoOmniitemItemImageUploadAPIRequest) GetImg() *model.File {
	return r._img
}

// Set is BarCode Setter
// 条形码
func (r *TaobaoOmniitemItemImageUploadAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// Get BarCode Getter
func (r TaobaoOmniitemItemImageUploadAPIRequest) GetBarCode() string {
	return r._barCode
}

// Set is ItemId Setter
// 商品ID，若填入商品ID则以商品ID为准，否则以outerId/barCode为准
func (r *TaobaoOmniitemItemImageUploadAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoOmniitemItemImageUploadAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is OuterId Setter
// 商品outerId
func (r *TaobaoOmniitemItemImageUploadAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TaobaoOmniitemItemImageUploadAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is Major Setter
// 是否为主图
func (r *TaobaoOmniitemItemImageUploadAPIRequest) SetMajor(_major bool) error {
	r._major = _major
	r.Set("major", _major)
	return nil
}

// Get Major Getter
func (r TaobaoOmniitemItemImageUploadAPIRequest) GetMajor() bool {
	return r._major
}

// Set is Position Setter
// 图片顺序
func (r *TaobaoOmniitemItemImageUploadAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// Get Position Getter
func (r TaobaoOmniitemItemImageUploadAPIRequest) GetPosition() int64 {
	return r._position
}
