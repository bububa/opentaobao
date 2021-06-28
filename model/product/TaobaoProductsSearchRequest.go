package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索产品信息 APIRequest
taobao.products.search

只有天猫商家发布商品时才需要用到，并非商品搜索api，当前暂不提供商品搜索api。<br/>二种方式搜索所有产品信息(二种至少传一种): <br/>
传入关键字q搜索<br/>
传入cid和props搜索<br/>
返回值支持:product_id,name,pic_path,cid,props,price,tsc<br/>
当用户指定了cid并且cid为垂直市场（3C电器城、鞋城）的类目id时，默认只返回小二确认<br/>的产品。如果用户没有指定cid，或cid为普通的类目，默认返回商家确认或小二确认的产<br/>品。如果用户自定了status字段，以指定的status类型为准。
<br/>新增一：
   传入suite_items_str 按规格搜索套装产品。
   返回字段增加suite_items_str,is_suite_effecitve支持。
*/
type TaobaoProductsSearchRequest struct {
    model.Params

    // 需返回的字段列表.可选值:Product数据结构中的以下字段:product_id,name,pic_url,cid,props,price,tsc;多个字段之间用","分隔.新增字段status(product的当前状态)
    fields   []string 

    // 搜索的关键词是用来搜索产品的title.　注:q,cid和props至少传入一个
    q   string 

    // 商品类目ID.调用taobao.itemcats.get获取.
    cid   int64 

    // 属性,属性值的组合.格式:pid:vid;pid:vid;调用taobao.itemprops.get获取类目属性pid <br/>,再用taobao.itempropvalues.get取得vid.
    props   string 

    // 想要获取的产品的状态列表，支持多个状态并列获取，多个状态之间用","分隔，最多同时指定5种状态。例如，只获取小二确认的spu传入"3",只要商家确认的传入"0"，既要小二确认又要商家确认的传入"0,3"。目前只支持者两种类型的状态搜索，输入其他状态无效。
    status   string 

    // 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始.
    pageNo   int64 

    // 每页条数.每页返回最多返回100条,默认值为40.
    pageSize   int64 

    // 传入值为：3表示3C表示3C垂直市场产品，4表示鞋城垂直市场产品，8表示网游垂直市场产品。一次只能指定一种垂直市场类型
    verticalMarket   int64 

    // 用户自定义关键属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”
    customerProps   string 

    // 按关联产品规格specs搜索套装产品
    suiteItemsStr   string 

    // 按条码搜索产品信息,多个逗号隔开，不支持条码为全零的方式
    barcodeStr   string 

    // 市场ID，1为取C2C市场的产品信息， 2为取B2C市场的产品信息。  不填写此值则默认取C2C的产品信息。
    marketId   string 

}

func NewTaobaoProductsSearchRequest() *TaobaoProductsSearchRequest{
    return &TaobaoProductsSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoProductsSearchRequest) GetApiMethodName() string {
    return "taobao.products.search"
}

func (r TaobaoProductsSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoProductsSearchRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoProductsSearchRequest) GetFields() []string {
    return r.fields
}

func (r *TaobaoProductsSearchRequest) SetQ(q string) error {
    r.q = q
    r.Set("q", q)
    return nil
}

func (r TaobaoProductsSearchRequest) GetQ() string {
    return r.q
}

func (r *TaobaoProductsSearchRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

func (r TaobaoProductsSearchRequest) GetCid() int64 {
    return r.cid
}

func (r *TaobaoProductsSearchRequest) SetProps(props string) error {
    r.props = props
    r.Set("props", props)
    return nil
}

func (r TaobaoProductsSearchRequest) GetProps() string {
    return r.props
}

func (r *TaobaoProductsSearchRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoProductsSearchRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoProductsSearchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoProductsSearchRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoProductsSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoProductsSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoProductsSearchRequest) SetVerticalMarket(verticalMarket int64) error {
    r.verticalMarket = verticalMarket
    r.Set("vertical_market", verticalMarket)
    return nil
}

func (r TaobaoProductsSearchRequest) GetVerticalMarket() int64 {
    return r.verticalMarket
}

func (r *TaobaoProductsSearchRequest) SetCustomerProps(customerProps string) error {
    r.customerProps = customerProps
    r.Set("customer_props", customerProps)
    return nil
}

func (r TaobaoProductsSearchRequest) GetCustomerProps() string {
    return r.customerProps
}

func (r *TaobaoProductsSearchRequest) SetSuiteItemsStr(suiteItemsStr string) error {
    r.suiteItemsStr = suiteItemsStr
    r.Set("suite_items_str", suiteItemsStr)
    return nil
}

func (r TaobaoProductsSearchRequest) GetSuiteItemsStr() string {
    return r.suiteItemsStr
}

func (r *TaobaoProductsSearchRequest) SetBarcodeStr(barcodeStr string) error {
    r.barcodeStr = barcodeStr
    r.Set("barcode_str", barcodeStr)
    return nil
}

func (r TaobaoProductsSearchRequest) GetBarcodeStr() string {
    return r.barcodeStr
}

func (r *TaobaoProductsSearchRequest) SetMarketId(marketId string) error {
    r.marketId = marketId
    r.Set("market_id", marketId)
    return nil
}

func (r TaobaoProductsSearchRequest) GetMarketId() string {
    return r.marketId
}

