package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家装新零售主商品同步至阿里 API请求
tmall.nrt.item.main.synchronize

同步红星美凯龙存量商品到阿里
*/
type TmallNrtItemMainSynchronizeRequest struct {
    model.Params
    // 摊位id
    boothId   string
    // 叶子类目id
    cid   int64
    // 类目属性
    props   []CategoryPropDto
    // 经销商编码
    dealerCode   string
    // 卖场id
    mallId   string
    // 商家编码
    outerId   string
    // 系统自动生成
    outerProps   *MacallineItemExtDto
    // 价格
    price   string
    // 商品名
    title   string
}

// 初始化TmallNrtItemMainSynchronizeRequest对象
func NewTmallNrtItemMainSynchronizeRequest() *TmallNrtItemMainSynchronizeRequest{
    return &TmallNrtItemMainSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtItemMainSynchronizeRequest) GetApiMethodName() string {
    return "tmall.nrt.item.main.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtItemMainSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BoothId Setter
// 摊位id
func (r *TmallNrtItemMainSynchronizeRequest) SetBoothId(boothId string) error {
    r.boothId = boothId
    r.Set("booth_id", boothId)
    return nil
}

// BoothId Getter
func (r TmallNrtItemMainSynchronizeRequest) GetBoothId() string {
    return r.boothId
}
// Cid Setter
// 叶子类目id
func (r *TmallNrtItemMainSynchronizeRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

// Cid Getter
func (r TmallNrtItemMainSynchronizeRequest) GetCid() int64 {
    return r.cid
}
// Props Setter
// 类目属性
func (r *TmallNrtItemMainSynchronizeRequest) SetProps(props []CategoryPropDto) error {
    r.props = props
    r.Set("props", props)
    return nil
}

// Props Getter
func (r TmallNrtItemMainSynchronizeRequest) GetProps() []CategoryPropDto {
    return r.props
}
// DealerCode Setter
// 经销商编码
func (r *TmallNrtItemMainSynchronizeRequest) SetDealerCode(dealerCode string) error {
    r.dealerCode = dealerCode
    r.Set("dealer_code", dealerCode)
    return nil
}

// DealerCode Getter
func (r TmallNrtItemMainSynchronizeRequest) GetDealerCode() string {
    return r.dealerCode
}
// MallId Setter
// 卖场id
func (r *TmallNrtItemMainSynchronizeRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

// MallId Getter
func (r TmallNrtItemMainSynchronizeRequest) GetMallId() string {
    return r.mallId
}
// OuterId Setter
// 商家编码
func (r *TmallNrtItemMainSynchronizeRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TmallNrtItemMainSynchronizeRequest) GetOuterId() string {
    return r.outerId
}
// OuterProps Setter
// 系统自动生成
func (r *TmallNrtItemMainSynchronizeRequest) SetOuterProps(outerProps *MacallineItemExtDto) error {
    r.outerProps = outerProps
    r.Set("outer_props", outerProps)
    return nil
}

// OuterProps Getter
func (r TmallNrtItemMainSynchronizeRequest) GetOuterProps() *MacallineItemExtDto {
    return r.outerProps
}
// Price Setter
// 价格
func (r *TmallNrtItemMainSynchronizeRequest) SetPrice(price string) error {
    r.price = price
    r.Set("price", price)
    return nil
}

// Price Getter
func (r TmallNrtItemMainSynchronizeRequest) GetPrice() string {
    return r.price
}
// Title Setter
// 商品名
func (r *TmallNrtItemMainSynchronizeRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TmallNrtItemMainSynchronizeRequest) GetTitle() string {
    return r.title
}
