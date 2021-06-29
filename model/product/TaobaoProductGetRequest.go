package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取一个产品的信息 API请求
taobao.product.get

天猫商家发布商品时，查询关联产品信息时使用，非商品查询接口。商品查询接口：taobao.item.seller.get<br>
两种方式查看一个产品详细信息: 
传入product_id来查询；传入cid和props来查询
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoProductGetRequest struct {
    model.Params
    // 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔.
    fields   string
    // Product的id.两种方式来查看一个产品:1.传入product_id来查询 2.传入cid和props来查询
    productId   int64
    // 商品类目id.调用taobao.itemcats.get获取;必须是叶子类目id,如果没有传product_id,那么cid和props必须要传.
    cid   int64
    // 比如:诺基亚N73这个产品的关键属性列表就是:品牌:诺基亚;型号:N73,对应的PV值就是10005:10027;10006:29729.
    props   string
}

// 初始化TaobaoProductGetRequest对象
func NewTaobaoProductGetRequest() *TaobaoProductGetRequest{
    return &TaobaoProductGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoProductGetRequest) GetApiMethodName() string {
    return "taobao.product.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoProductGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔.
func (r *TaobaoProductGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoProductGetRequest) GetFields() string {
    return r.fields
}
// ProductId Setter
// Product的id.两种方式来查看一个产品:1.传入product_id来查询 2.传入cid和props来查询
func (r *TaobaoProductGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoProductGetRequest) GetProductId() int64 {
    return r.productId
}
// Cid Setter
// 商品类目id.调用taobao.itemcats.get获取;必须是叶子类目id,如果没有传product_id,那么cid和props必须要传.
func (r *TaobaoProductGetRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

// Cid Getter
func (r TaobaoProductGetRequest) GetCid() int64 {
    return r.cid
}
// Props Setter
// 比如:诺基亚N73这个产品的关键属性列表就是:品牌:诺基亚;型号:N73,对应的PV值就是10005:10027;10006:29729.
func (r *TaobaoProductGetRequest) SetProps(props string) error {
    r.props = props
    r.Set("props", props)
    return nil
}

// Props Getter
func (r TaobaoProductGetRequest) GetProps() string {
    return r.props
}
