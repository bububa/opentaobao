package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传一个产品，不包括产品非主图和属性图片 APIRequest
taobao.product.add

获取类目ID，必需是叶子类目ID；调用taobao.itemcats.get.v2获取 <br/>传入关键属性,结构:pid:vid;pid:vid.调用taobao.itemprops.get.v2获取pid,<br/>调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.<br/>新增：套装产品发布，目前支持单件多个即 A*2 形式的套装
*/
type TaobaoProductAddRequest struct {
    model.Params

    // native_unkeyprops
    nativeUnkeyprops   string 

    // 商品类目ID.调用taobao.itemcats.get获取;注意:必须是叶子类目 id.
    cid   int64 

    // 外部产品ID
    outerId   string 

    // 关键属性 结构:pid:vid;pid:vid.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.
    props   string 

    // 非关键属性结构:pid:vid;pid:vid.<br>
非关键属性<font color=red>不包含</font>关键属性、销售属性、用户自定义属性、商品属性;
<br>调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid.<br><font color=red>注:支持最大长度为512字节</font>
    binds   string 

    // 销售属性结构:pid:vid;pid:vid.调用taobao.itemprops.get获取is_sale_prop＝true的pid,调用taobao.itempropvalues.get获取vid.
    saleProps   string 

    // 用户自定义属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”
<br><font color=red>注：包含所有自定义属性的传入</font>
    customerProps   string 

    // 产品市场价.精确到2位小数;单位为元.如：200.07
    price   string 

    // 产品主图片.最大1M,目前仅支持GIF,JPG.
    image   []byte 

    // 产品名称,最大30个字符.
    name   string 

    // 产品描述.最大不超过25000个字符
    desc   string 

    // 是不是主图
    major   bool 

    // 上市时间。目前只支持鞋城类目传入此参数
    marketTime   string 

    // 销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。只有少数销售属性值支持传入别名，比如颜色和尺寸
    propertyAlias   string 

}

func NewTaobaoProductAddRequest() *TaobaoProductAddRequest{
    return &TaobaoProductAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoProductAddRequest) GetApiMethodName() string {
    return "taobao.product.add"
}

func (r TaobaoProductAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoProductAddRequest) SetNativeUnkeyprops(nativeUnkeyprops string) error {
    r.nativeUnkeyprops = nativeUnkeyprops
    r.Set("native_unkeyprops", nativeUnkeyprops)
    return nil
}

func (r TaobaoProductAddRequest) GetNativeUnkeyprops() string {
    return r.nativeUnkeyprops
}

func (r *TaobaoProductAddRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

func (r TaobaoProductAddRequest) GetCid() int64 {
    return r.cid
}

func (r *TaobaoProductAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoProductAddRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoProductAddRequest) SetProps(props string) error {
    r.props = props
    r.Set("props", props)
    return nil
}

func (r TaobaoProductAddRequest) GetProps() string {
    return r.props
}

func (r *TaobaoProductAddRequest) SetBinds(binds string) error {
    r.binds = binds
    r.Set("binds", binds)
    return nil
}

func (r TaobaoProductAddRequest) GetBinds() string {
    return r.binds
}

func (r *TaobaoProductAddRequest) SetSaleProps(saleProps string) error {
    r.saleProps = saleProps
    r.Set("sale_props", saleProps)
    return nil
}

func (r TaobaoProductAddRequest) GetSaleProps() string {
    return r.saleProps
}

func (r *TaobaoProductAddRequest) SetCustomerProps(customerProps string) error {
    r.customerProps = customerProps
    r.Set("customer_props", customerProps)
    return nil
}

func (r TaobaoProductAddRequest) GetCustomerProps() string {
    return r.customerProps
}

func (r *TaobaoProductAddRequest) SetPrice(price string) error {
    r.price = price
    r.Set("price", price)
    return nil
}

func (r TaobaoProductAddRequest) GetPrice() string {
    return r.price
}

func (r *TaobaoProductAddRequest) SetImage(image []byte) error {
    r.image = image
    r.Set("image", image)
    return nil
}

func (r TaobaoProductAddRequest) GetImage() []byte {
    return r.image
}

func (r *TaobaoProductAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoProductAddRequest) GetName() string {
    return r.name
}

func (r *TaobaoProductAddRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r TaobaoProductAddRequest) GetDesc() string {
    return r.desc
}

func (r *TaobaoProductAddRequest) SetMajor(major bool) error {
    r.major = major
    r.Set("major", major)
    return nil
}

func (r TaobaoProductAddRequest) GetMajor() bool {
    return r.major
}

func (r *TaobaoProductAddRequest) SetMarketTime(marketTime string) error {
    r.marketTime = marketTime
    r.Set("market_time", marketTime)
    return nil
}

func (r TaobaoProductAddRequest) GetMarketTime() string {
    return r.marketTime
}

func (r *TaobaoProductAddRequest) SetPropertyAlias(propertyAlias string) error {
    r.propertyAlias = propertyAlias
    r.Set("property_alias", propertyAlias)
    return nil
}

func (r TaobaoProductAddRequest) GetPropertyAlias() string {
    return r.propertyAlias
}

