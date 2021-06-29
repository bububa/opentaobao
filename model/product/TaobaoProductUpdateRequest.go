package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改一个产品，可以修改主图，不能修改子图片 API请求
taobao.product.update

传入产品ID <br/>可修改字段：outer_id,binds,sale_props,name,price,desc,image <br/>注意：1.可以修改主图,不能修改子图片,主图最大500K,目前仅支持GIF,JPG<br/>      2.商城卖家产品发布24小时后不能作删除或修改操作
*/
type TaobaoProductUpdateRequest struct {
    model.Params
    // 产品ID
    productId   int64
    // 外部产品ID
    outerId   string
    // 非关键属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
    binds   string
    // 销售属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
    saleProps   string
    // 产品市场价.精确到2位小数;单位为元.如:200.07
    price   string
    // 产品描述.最大不超过25000个字符
    desc   string
    // 产品主图.最大500K,目前仅支持GIF,JPG
    image   []*model.File
    // 产品名称.最大不超过30个字符
    name   string
    // 是否是主图
    major   bool
    // 自定义非关键属性
    nativeUnkeyprops   string
}

// 初始化TaobaoProductUpdateRequest对象
func NewTaobaoProductUpdateRequest() *TaobaoProductUpdateRequest{
    return &TaobaoProductUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoProductUpdateRequest) GetApiMethodName() string {
    return "taobao.product.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoProductUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TaobaoProductUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoProductUpdateRequest) GetProductId() int64 {
    return r.productId
}
// OuterId Setter
// 外部产品ID
func (r *TaobaoProductUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoProductUpdateRequest) GetOuterId() string {
    return r.outerId
}
// Binds Setter
// 非关键属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
func (r *TaobaoProductUpdateRequest) SetBinds(binds string) error {
    r.binds = binds
    r.Set("binds", binds)
    return nil
}

// Binds Getter
func (r TaobaoProductUpdateRequest) GetBinds() string {
    return r.binds
}
// SaleProps Setter
// 销售属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
func (r *TaobaoProductUpdateRequest) SetSaleProps(saleProps string) error {
    r.saleProps = saleProps
    r.Set("sale_props", saleProps)
    return nil
}

// SaleProps Getter
func (r TaobaoProductUpdateRequest) GetSaleProps() string {
    return r.saleProps
}
// Price Setter
// 产品市场价.精确到2位小数;单位为元.如:200.07
func (r *TaobaoProductUpdateRequest) SetPrice(price string) error {
    r.price = price
    r.Set("price", price)
    return nil
}

// Price Getter
func (r TaobaoProductUpdateRequest) GetPrice() string {
    return r.price
}
// Desc Setter
// 产品描述.最大不超过25000个字符
func (r *TaobaoProductUpdateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

// Desc Getter
func (r TaobaoProductUpdateRequest) GetDesc() string {
    return r.desc
}
// Image Setter
// 产品主图.最大500K,目前仅支持GIF,JPG
func (r *TaobaoProductUpdateRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TaobaoProductUpdateRequest) GetImage() []*model.File {
    return r.image
}
// Name Setter
// 产品名称.最大不超过30个字符
func (r *TaobaoProductUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoProductUpdateRequest) GetName() string {
    return r.name
}
// Major Setter
// 是否是主图
func (r *TaobaoProductUpdateRequest) SetMajor(major bool) error {
    r.major = major
    r.Set("major", major)
    return nil
}

// Major Getter
func (r TaobaoProductUpdateRequest) GetMajor() bool {
    return r.major
}
// NativeUnkeyprops Setter
// 自定义非关键属性
func (r *TaobaoProductUpdateRequest) SetNativeUnkeyprops(nativeUnkeyprops string) error {
    r.nativeUnkeyprops = nativeUnkeyprops
    r.Set("native_unkeyprops", nativeUnkeyprops)
    return nil
}

// NativeUnkeyprops Getter
func (r TaobaoProductUpdateRequest) GetNativeUnkeyprops() string {
    return r.nativeUnkeyprops
}
