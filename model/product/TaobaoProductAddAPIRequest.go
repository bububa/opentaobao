package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductAddAPIRequest 上传一个产品，不包括产品非主图和属性图片 API请求
// taobao.product.add
//
// 获取类目ID，必需是叶子类目ID；调用taobao.itemcats.get.v2获取 &lt;br/&gt;传入关键属性,结构:pid:vid;pid:vid.调用taobao.itemprops.get.v2获取pid,&lt;br/&gt;调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.&lt;br/&gt;新增：套装产品发布，目前支持单件多个即 A*2 形式的套装
type TaobaoProductAddAPIRequest struct {
	model.Params
	// 产品名称,最大30个字符.
	_name string
	// 产品市场价.精确到2位小数;单位为元.如：200.07
	_price string
	// 外部产品ID
	_outerId string
	// 关键属性 结构:pid:vid;pid:vid.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.
	_props string
	// 非关键属性结构:pid:vid;pid:vid.<br>非关键属性<font color=red>不包含</font>关键属性、销售属性、用户自定义属性、商品属性;<br>调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid.<br><font color=red>注:支持最大长度为512字节</font>
	_binds string
	// 销售属性结构:pid:vid;pid:vid.调用taobao.itemprops.get获取is_sale_prop＝true的pid,调用taobao.itempropvalues.get获取vid.
	_saleProps string
	// 用户自定义属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”<br><font color=red>注：包含所有自定义属性的传入</font>
	_customerProps string
	// 产品描述.最大不超过25000个字符
	_desc string
	// native_unkeyprops
	_nativeUnkeyprops string
	// 上市时间。目前只支持鞋城类目传入此参数
	_marketTime string
	// 销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。只有少数销售属性值支持传入别名，比如颜色和尺寸
	_propertyAlias string
	// 商品类目ID.调用taobao.itemcats.get获取;注意:必须是叶子类目 id.
	_cid int64
	// 产品主图片.最大1M,目前仅支持GIF,JPG.
	_image *model.File
	// 是不是主图
	_major bool
}

// NewTaobaoProductAddRequest 初始化TaobaoProductAddAPIRequest对象
func NewTaobaoProductAddRequest() *TaobaoProductAddAPIRequest {
	return &TaobaoProductAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoProductAddAPIRequest) GetApiMethodName() string {
	return "taobao.product.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoProductAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetName is Name Setter
// 产品名称,最大30个字符.
func (r *TaobaoProductAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoProductAddAPIRequest) GetName() string {
	return r._name
}

// SetPrice is Price Setter
// 产品市场价.精确到2位小数;单位为元.如：200.07
func (r *TaobaoProductAddAPIRequest) SetPrice(_price string) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TaobaoProductAddAPIRequest) GetPrice() string {
	return r._price
}

// SetOuterId is OuterId Setter
// 外部产品ID
func (r *TaobaoProductAddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoProductAddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetProps is Props Setter
// 关键属性 结构:pid:vid;pid:vid.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.
func (r *TaobaoProductAddAPIRequest) SetProps(_props string) error {
	r._props = _props
	r.Set("props", _props)
	return nil
}

// GetProps Props Getter
func (r TaobaoProductAddAPIRequest) GetProps() string {
	return r._props
}

// SetBinds is Binds Setter
// 非关键属性结构:pid:vid;pid:vid.&lt;br&gt;非关键属性&lt;font color=red&gt;不包含&lt;/font&gt;关键属性、销售属性、用户自定义属性、商品属性;&lt;br&gt;调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid.&lt;br&gt;&lt;font color=red&gt;注:支持最大长度为512字节&lt;/font&gt;
func (r *TaobaoProductAddAPIRequest) SetBinds(_binds string) error {
	r._binds = _binds
	r.Set("binds", _binds)
	return nil
}

// GetBinds Binds Getter
func (r TaobaoProductAddAPIRequest) GetBinds() string {
	return r._binds
}

// SetSaleProps is SaleProps Setter
// 销售属性结构:pid:vid;pid:vid.调用taobao.itemprops.get获取is_sale_prop＝true的pid,调用taobao.itempropvalues.get获取vid.
func (r *TaobaoProductAddAPIRequest) SetSaleProps(_saleProps string) error {
	r._saleProps = _saleProps
	r.Set("sale_props", _saleProps)
	return nil
}

// GetSaleProps SaleProps Getter
func (r TaobaoProductAddAPIRequest) GetSaleProps() string {
	return r._saleProps
}

// SetCustomerProps is CustomerProps Setter
// 用户自定义属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”&lt;br&gt;&lt;font color=red&gt;注：包含所有自定义属性的传入&lt;/font&gt;
func (r *TaobaoProductAddAPIRequest) SetCustomerProps(_customerProps string) error {
	r._customerProps = _customerProps
	r.Set("customer_props", _customerProps)
	return nil
}

// GetCustomerProps CustomerProps Getter
func (r TaobaoProductAddAPIRequest) GetCustomerProps() string {
	return r._customerProps
}

// SetDesc is Desc Setter
// 产品描述.最大不超过25000个字符
func (r *TaobaoProductAddAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoProductAddAPIRequest) GetDesc() string {
	return r._desc
}

// SetNativeUnkeyprops is NativeUnkeyprops Setter
// native_unkeyprops
func (r *TaobaoProductAddAPIRequest) SetNativeUnkeyprops(_nativeUnkeyprops string) error {
	r._nativeUnkeyprops = _nativeUnkeyprops
	r.Set("native_unkeyprops", _nativeUnkeyprops)
	return nil
}

// GetNativeUnkeyprops NativeUnkeyprops Getter
func (r TaobaoProductAddAPIRequest) GetNativeUnkeyprops() string {
	return r._nativeUnkeyprops
}

// SetMarketTime is MarketTime Setter
// 上市时间。目前只支持鞋城类目传入此参数
func (r *TaobaoProductAddAPIRequest) SetMarketTime(_marketTime string) error {
	r._marketTime = _marketTime
	r.Set("market_time", _marketTime)
	return nil
}

// GetMarketTime MarketTime Getter
func (r TaobaoProductAddAPIRequest) GetMarketTime() string {
	return r._marketTime
}

// SetPropertyAlias is PropertyAlias Setter
// 销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。只有少数销售属性值支持传入别名，比如颜色和尺寸
func (r *TaobaoProductAddAPIRequest) SetPropertyAlias(_propertyAlias string) error {
	r._propertyAlias = _propertyAlias
	r.Set("property_alias", _propertyAlias)
	return nil
}

// GetPropertyAlias PropertyAlias Getter
func (r TaobaoProductAddAPIRequest) GetPropertyAlias() string {
	return r._propertyAlias
}

// SetCid is Cid Setter
// 商品类目ID.调用taobao.itemcats.get获取;注意:必须是叶子类目 id.
func (r *TaobaoProductAddAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoProductAddAPIRequest) GetCid() int64 {
	return r._cid
}

// SetImage is Image Setter
// 产品主图片.最大1M,目前仅支持GIF,JPG.
func (r *TaobaoProductAddAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoProductAddAPIRequest) GetImage() *model.File {
	return r._image
}

// SetMajor is Major Setter
// 是不是主图
func (r *TaobaoProductAddAPIRequest) SetMajor(_major bool) error {
	r._major = _major
	r.Set("major", _major)
	return nil
}

// GetMajor Major Getter
func (r TaobaoProductAddAPIRequest) GetMajor() bool {
	return r._major
}
