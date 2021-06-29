package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传一个产品，不包括产品非主图和属性图片 API请求
taobao.product.add

获取类目ID，必需是叶子类目ID；调用taobao.itemcats.get.v2获取 <br/>传入关键属性,结构:pid:vid;pid:vid.调用taobao.itemprops.get.v2获取pid,<br/>调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.<br/>新增：套装产品发布，目前支持单件多个即 A*2 形式的套装
*/
type TaobaoProductAddRequest struct {
    model.Params
    // native_unkeyprops
    _nativeUnkeyprops   string
    // 商品类目ID.调用taobao.itemcats.get获取;注意:必须是叶子类目 id.
    _cid   int64
    // 外部产品ID
    _outerId   string
    // 关键属性 结构:pid:vid;pid:vid.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.
    _props   string
    // 非关键属性结构:pid:vid;pid:vid.<br>非关键属性<font color=red>不包含</font>关键属性、销售属性、用户自定义属性、商品属性;<br>调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid.<br><font color=red>注:支持最大长度为512字节</font>
    _binds   string
    // 销售属性结构:pid:vid;pid:vid.调用taobao.itemprops.get获取is_sale_prop＝true的pid,调用taobao.itempropvalues.get获取vid.
    _saleProps   string
    // 用户自定义属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”<br><font color=red>注：包含所有自定义属性的传入</font>
    _customerProps   string
    // 产品市场价.精确到2位小数;单位为元.如：200.07
    _price   string
    // 产品主图片.最大1M,目前仅支持GIF,JPG.
    _image   *model.File
    // 产品名称,最大30个字符.
    _name   string
    // 产品描述.最大不超过25000个字符
    _desc   string
    // 是不是主图
    _major   bool
    // 上市时间。目前只支持鞋城类目传入此参数
    _marketTime   string
    // 销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。只有少数销售属性值支持传入别名，比如颜色和尺寸
    _propertyAlias   string
}

// 初始化TaobaoProductAddRequest对象
func NewTaobaoProductAddRequest() *TaobaoProductAddRequest{
    return &TaobaoProductAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoProductAddRequest) GetApiMethodName() string {
    return "taobao.product.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoProductAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NativeUnkeyprops Setter
// native_unkeyprops
func (r *TaobaoProductAddRequest) SetNativeUnkeyprops(_nativeUnkeyprops string) error {
    r._nativeUnkeyprops = _nativeUnkeyprops
    r.Set("native_unkeyprops", _nativeUnkeyprops)
    return nil
}

// NativeUnkeyprops Getter
func (r TaobaoProductAddRequest) GetNativeUnkeyprops() string {
    return r._nativeUnkeyprops
}
// Cid Setter
// 商品类目ID.调用taobao.itemcats.get获取;注意:必须是叶子类目 id.
func (r *TaobaoProductAddRequest) SetCid(_cid int64) error {
    r._cid = _cid
    r.Set("cid", _cid)
    return nil
}

// Cid Getter
func (r TaobaoProductAddRequest) GetCid() int64 {
    return r._cid
}
// OuterId Setter
// 外部产品ID
func (r *TaobaoProductAddRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoProductAddRequest) GetOuterId() string {
    return r._outerId
}
// Props Setter
// 关键属性 结构:pid:vid;pid:vid.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.
func (r *TaobaoProductAddRequest) SetProps(_props string) error {
    r._props = _props
    r.Set("props", _props)
    return nil
}

// Props Getter
func (r TaobaoProductAddRequest) GetProps() string {
    return r._props
}
// Binds Setter
// 非关键属性结构:pid:vid;pid:vid.<br>非关键属性<font color=red>不包含</font>关键属性、销售属性、用户自定义属性、商品属性;<br>调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid.<br><font color=red>注:支持最大长度为512字节</font>
func (r *TaobaoProductAddRequest) SetBinds(_binds string) error {
    r._binds = _binds
    r.Set("binds", _binds)
    return nil
}

// Binds Getter
func (r TaobaoProductAddRequest) GetBinds() string {
    return r._binds
}
// SaleProps Setter
// 销售属性结构:pid:vid;pid:vid.调用taobao.itemprops.get获取is_sale_prop＝true的pid,调用taobao.itempropvalues.get获取vid.
func (r *TaobaoProductAddRequest) SetSaleProps(_saleProps string) error {
    r._saleProps = _saleProps
    r.Set("sale_props", _saleProps)
    return nil
}

// SaleProps Getter
func (r TaobaoProductAddRequest) GetSaleProps() string {
    return r._saleProps
}
// CustomerProps Setter
// 用户自定义属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”<br><font color=red>注：包含所有自定义属性的传入</font>
func (r *TaobaoProductAddRequest) SetCustomerProps(_customerProps string) error {
    r._customerProps = _customerProps
    r.Set("customer_props", _customerProps)
    return nil
}

// CustomerProps Getter
func (r TaobaoProductAddRequest) GetCustomerProps() string {
    return r._customerProps
}
// Price Setter
// 产品市场价.精确到2位小数;单位为元.如：200.07
func (r *TaobaoProductAddRequest) SetPrice(_price string) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r TaobaoProductAddRequest) GetPrice() string {
    return r._price
}
// Image Setter
// 产品主图片.最大1M,目前仅支持GIF,JPG.
func (r *TaobaoProductAddRequest) SetImage(_image *model.File) error {
    r._image = _image
    r.Set("image", _image)
    return nil
}

// Image Getter
func (r TaobaoProductAddRequest) GetImage() *model.File {
    return r._image
}
// Name Setter
// 产品名称,最大30个字符.
func (r *TaobaoProductAddRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoProductAddRequest) GetName() string {
    return r._name
}
// Desc Setter
// 产品描述.最大不超过25000个字符
func (r *TaobaoProductAddRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r TaobaoProductAddRequest) GetDesc() string {
    return r._desc
}
// Major Setter
// 是不是主图
func (r *TaobaoProductAddRequest) SetMajor(_major bool) error {
    r._major = _major
    r.Set("major", _major)
    return nil
}

// Major Getter
func (r TaobaoProductAddRequest) GetMajor() bool {
    return r._major
}
// MarketTime Setter
// 上市时间。目前只支持鞋城类目传入此参数
func (r *TaobaoProductAddRequest) SetMarketTime(_marketTime string) error {
    r._marketTime = _marketTime
    r.Set("market_time", _marketTime)
    return nil
}

// MarketTime Getter
func (r TaobaoProductAddRequest) GetMarketTime() string {
    return r._marketTime
}
// PropertyAlias Setter
// 销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。只有少数销售属性值支持传入别名，比如颜色和尺寸
func (r *TaobaoProductAddRequest) SetPropertyAlias(_propertyAlias string) error {
    r._propertyAlias = _propertyAlias
    r.Set("property_alias", _propertyAlias)
    return nil
}

// PropertyAlias Getter
func (r TaobaoProductAddRequest) GetPropertyAlias() string {
    return r._propertyAlias
}
