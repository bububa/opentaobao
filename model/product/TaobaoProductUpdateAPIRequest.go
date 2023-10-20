package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductUpdateAPIRequest 修改一个产品，可以修改主图，不能修改子图片 API请求
// taobao.product.update
//
// 传入产品ID &lt;br/&gt;可修改字段：outer_id,binds,sale_props,name,price,desc,image &lt;br/&gt;注意：1.可以修改主图,不能修改子图片,主图最大500K,目前仅支持GIF,JPG&lt;br/&gt;      2.商城卖家产品发布24小时后不能作删除或修改操作
type TaobaoProductUpdateAPIRequest struct {
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

// NewTaobaoProductUpdateRequest 初始化TaobaoProductUpdateAPIRequest对象
func NewTaobaoProductUpdateRequest() *TaobaoProductUpdateAPIRequest {
	return &TaobaoProductUpdateAPIRequest{
		Params: model.NewParams(10),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoProductUpdateAPIRequest) Reset() {
	r._outerId = ""
	r._binds = ""
	r._saleProps = ""
	r._name = ""
	r._price = ""
	r._desc = ""
	r._nativeUnkeyprops = ""
	r._productId = 0
	r._image = nil
	r._major = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoProductUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.product.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoProductUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoProductUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部产品ID
func (r *TaobaoProductUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoProductUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetBinds is Binds Setter
// 非关键属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
func (r *TaobaoProductUpdateAPIRequest) SetBinds(_binds string) error {
	r._binds = _binds
	r.Set("binds", _binds)
	return nil
}

// GetBinds Binds Getter
func (r TaobaoProductUpdateAPIRequest) GetBinds() string {
	return r._binds
}

// SetSaleProps is SaleProps Setter
// 销售属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
func (r *TaobaoProductUpdateAPIRequest) SetSaleProps(_saleProps string) error {
	r._saleProps = _saleProps
	r.Set("sale_props", _saleProps)
	return nil
}

// GetSaleProps SaleProps Getter
func (r TaobaoProductUpdateAPIRequest) GetSaleProps() string {
	return r._saleProps
}

// SetName is Name Setter
// 产品名称.最大不超过30个字符
func (r *TaobaoProductUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoProductUpdateAPIRequest) GetName() string {
	return r._name
}

// SetPrice is Price Setter
// 产品市场价.精确到2位小数;单位为元.如:200.07
func (r *TaobaoProductUpdateAPIRequest) SetPrice(_price string) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TaobaoProductUpdateAPIRequest) GetPrice() string {
	return r._price
}

// SetDesc is Desc Setter
// 产品描述.最大不超过25000个字符
func (r *TaobaoProductUpdateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoProductUpdateAPIRequest) GetDesc() string {
	return r._desc
}

// SetNativeUnkeyprops is NativeUnkeyprops Setter
// 自定义非关键属性
func (r *TaobaoProductUpdateAPIRequest) SetNativeUnkeyprops(_nativeUnkeyprops string) error {
	r._nativeUnkeyprops = _nativeUnkeyprops
	r.Set("native_unkeyprops", _nativeUnkeyprops)
	return nil
}

// GetNativeUnkeyprops NativeUnkeyprops Getter
func (r TaobaoProductUpdateAPIRequest) GetNativeUnkeyprops() string {
	return r._nativeUnkeyprops
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaoProductUpdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoProductUpdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetImage is Image Setter
// 产品主图.最大500K,目前仅支持GIF,JPG
func (r *TaobaoProductUpdateAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoProductUpdateAPIRequest) GetImage() *model.File {
	return r._image
}

// SetMajor is Major Setter
// 是否是主图
func (r *TaobaoProductUpdateAPIRequest) SetMajor(_major bool) error {
	r._major = _major
	r.Set("major", _major)
	return nil
}

// GetMajor Major Getter
func (r TaobaoProductUpdateAPIRequest) GetMajor() bool {
	return r._major
}

var poolTaobaoProductUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoProductUpdateRequest()
	},
}

// GetTaobaoProductUpdateRequest 从 sync.Pool 获取 TaobaoProductUpdateAPIRequest
func GetTaobaoProductUpdateAPIRequest() *TaobaoProductUpdateAPIRequest {
	return poolTaobaoProductUpdateAPIRequest.Get().(*TaobaoProductUpdateAPIRequest)
}

// ReleaseTaobaoProductUpdateAPIRequest 将 TaobaoProductUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoProductUpdateAPIRequest(v *TaobaoProductUpdateAPIRequest) {
	v.Reset()
	poolTaobaoProductUpdateAPIRequest.Put(v)
}
