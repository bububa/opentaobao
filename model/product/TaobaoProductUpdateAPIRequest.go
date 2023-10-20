package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoproductupdateAPIRequest 修改一个产品，可以修改主图，不能修改子图片 API请求
// taobao.product.update
//
// 传入产品ID &lt;br/&gt;可修改字段：outer_id,binds,sale_props,name,price,desc,image &lt;br/&gt;注意：1.可以修改主图,不能修改子图片,主图最大500K,目前仅支持GIF,JPG&lt;br/&gt;      2.商城卖家产品发布24小时后不能作删除或修改操作
type TaobaoproductupdateAPIRequest struct {
	model.Params
	// 外部产品ID
	_outerId string
	// 非关键属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
	_binds string
	// 销售属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
	_saleProps string
	// 产品名称.最大不超过30个字符
	_name string
	// 产品市场价.精确到2位小数;单位为元.如:200.07
	_price string
	// 产品描述.最大不超过25000个字符
	_desc string
	// 自定义非关键属性
	_nativeUnkeyprops string
	// 产品ID
	_productId int64
	// 产品主图.最大500K,目前仅支持GIF,JPG
	_image *model.File
	// 是否是主图
	_major bool
}

// NewTaobaoproductupdateRequest 初始化TaobaoproductupdateAPIRequest对象
func NewTaobaoproductupdateRequest() *TaobaoproductupdateAPIRequest {
	return &TaobaoproductupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoproductupdateAPIRequest) GetApiMethodName() string {
	return "taobao.product.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoproductupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoproductupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部产品ID
func (r *TaobaoproductupdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoproductupdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetBinds is Binds Setter
// 非关键属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
func (r *TaobaoproductupdateAPIRequest) SetBinds(_binds string) error {
	r._binds = _binds
	r.Set("binds", _binds)
	return nil
}

// GetBinds Binds Getter
func (r TaobaoproductupdateAPIRequest) GetBinds() string {
	return r._binds
}

// SetSaleProps is SaleProps Setter
// 销售属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
func (r *TaobaoproductupdateAPIRequest) SetSaleProps(_saleProps string) error {
	r._saleProps = _saleProps
	r.Set("sale_props", _saleProps)
	return nil
}

// GetSaleProps SaleProps Getter
func (r TaobaoproductupdateAPIRequest) GetSaleProps() string {
	return r._saleProps
}

// SetName is Name Setter
// 产品名称.最大不超过30个字符
func (r *TaobaoproductupdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoproductupdateAPIRequest) GetName() string {
	return r._name
}

// SetPrice is Price Setter
// 产品市场价.精确到2位小数;单位为元.如:200.07
func (r *TaobaoproductupdateAPIRequest) SetPrice(_price string) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TaobaoproductupdateAPIRequest) GetPrice() string {
	return r._price
}

// SetDesc is Desc Setter
// 产品描述.最大不超过25000个字符
func (r *TaobaoproductupdateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoproductupdateAPIRequest) GetDesc() string {
	return r._desc
}

// SetNativeUnkeyprops is NativeUnkeyprops Setter
// 自定义非关键属性
func (r *TaobaoproductupdateAPIRequest) SetNativeUnkeyprops(_nativeUnkeyprops string) error {
	r._nativeUnkeyprops = _nativeUnkeyprops
	r.Set("native_unkeyprops", _nativeUnkeyprops)
	return nil
}

// GetNativeUnkeyprops NativeUnkeyprops Getter
func (r TaobaoproductupdateAPIRequest) GetNativeUnkeyprops() string {
	return r._nativeUnkeyprops
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaoproductupdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoproductupdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetImage is Image Setter
// 产品主图.最大500K,目前仅支持GIF,JPG
func (r *TaobaoproductupdateAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoproductupdateAPIRequest) GetImage() *model.File {
	return r._image
}

// SetMajor is Major Setter
// 是否是主图
func (r *TaobaoproductupdateAPIRequest) SetMajor(_major bool) error {
	r._major = _major
	r.Set("major", _major)
	return nil
}

// GetMajor Major Getter
func (r TaobaoproductupdateAPIRequest) GetMajor() bool {
	return r._major
}
