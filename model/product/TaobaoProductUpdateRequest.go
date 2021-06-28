package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改一个产品，可以修改主图，不能修改子图片 APIRequest
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

func NewTaobaoProductUpdateRequest() *TaobaoProductUpdateRequest{
    return &TaobaoProductUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoProductUpdateRequest) GetApiMethodName() string {
    return "taobao.product.update"
}

func (r TaobaoProductUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoProductUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoProductUpdateRequest) GetProductId() int64 {
    return r.productId
}

func (r *TaobaoProductUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoProductUpdateRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoProductUpdateRequest) SetBinds(binds string) error {
    r.binds = binds
    r.Set("binds", binds)
    return nil
}

func (r TaobaoProductUpdateRequest) GetBinds() string {
    return r.binds
}

func (r *TaobaoProductUpdateRequest) SetSaleProps(saleProps string) error {
    r.saleProps = saleProps
    r.Set("sale_props", saleProps)
    return nil
}

func (r TaobaoProductUpdateRequest) GetSaleProps() string {
    return r.saleProps
}

func (r *TaobaoProductUpdateRequest) SetPrice(price string) error {
    r.price = price
    r.Set("price", price)
    return nil
}

func (r TaobaoProductUpdateRequest) GetPrice() string {
    return r.price
}

func (r *TaobaoProductUpdateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r TaobaoProductUpdateRequest) GetDesc() string {
    return r.desc
}

func (r *TaobaoProductUpdateRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

func (r TaobaoProductUpdateRequest) GetImage() []*model.File {
    return r.image
}

func (r *TaobaoProductUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoProductUpdateRequest) GetName() string {
    return r.name
}

func (r *TaobaoProductUpdateRequest) SetMajor(major bool) error {
    r.major = major
    r.Set("major", major)
    return nil
}

func (r TaobaoProductUpdateRequest) GetMajor() bool {
    return r.major
}

func (r *TaobaoProductUpdateRequest) SetNativeUnkeyprops(nativeUnkeyprops string) error {
    r.nativeUnkeyprops = nativeUnkeyprops
    r.Set("native_unkeyprops", nativeUnkeyprops)
    return nil
}

func (r TaobaoProductUpdateRequest) GetNativeUnkeyprops() string {
    return r.nativeUnkeyprops
}

