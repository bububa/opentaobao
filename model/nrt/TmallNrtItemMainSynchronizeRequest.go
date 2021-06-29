package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家装新零售主商品同步至阿里 APIRequest
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

func NewTmallNrtItemMainSynchronizeRequest() *TmallNrtItemMainSynchronizeRequest{
    return &TmallNrtItemMainSynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtItemMainSynchronizeRequest) GetApiMethodName() string {
    return "tmall.nrt.item.main.synchronize"
}

func (r TmallNrtItemMainSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtItemMainSynchronizeRequest) SetBoothId(boothId string) error {
    r.boothId = boothId
    r.Set("booth_id", boothId)
    return nil
}

func (r TmallNrtItemMainSynchronizeRequest) GetBoothId() string {
    return r.boothId
}

func (r *TmallNrtItemMainSynchronizeRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

func (r TmallNrtItemMainSynchronizeRequest) GetCid() int64 {
    return r.cid
}

func (r *TmallNrtItemMainSynchronizeRequest) SetProps(props []CategoryPropDto) error {
    r.props = props
    r.Set("props", props)
    return nil
}

func (r TmallNrtItemMainSynchronizeRequest) GetProps() []CategoryPropDto {
    return r.props
}

func (r *TmallNrtItemMainSynchronizeRequest) SetDealerCode(dealerCode string) error {
    r.dealerCode = dealerCode
    r.Set("dealer_code", dealerCode)
    return nil
}

func (r TmallNrtItemMainSynchronizeRequest) GetDealerCode() string {
    return r.dealerCode
}

func (r *TmallNrtItemMainSynchronizeRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r TmallNrtItemMainSynchronizeRequest) GetMallId() string {
    return r.mallId
}

func (r *TmallNrtItemMainSynchronizeRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TmallNrtItemMainSynchronizeRequest) GetOuterId() string {
    return r.outerId
}

func (r *TmallNrtItemMainSynchronizeRequest) SetOuterProps(outerProps *MacallineItemExtDto) error {
    r.outerProps = outerProps
    r.Set("outer_props", outerProps)
    return nil
}

func (r TmallNrtItemMainSynchronizeRequest) GetOuterProps() *MacallineItemExtDto {
    return r.outerProps
}

func (r *TmallNrtItemMainSynchronizeRequest) SetPrice(price string) error {
    r.price = price
    r.Set("price", price)
    return nil
}

func (r TmallNrtItemMainSynchronizeRequest) GetPrice() string {
    return r.price
}

func (r *TmallNrtItemMainSynchronizeRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TmallNrtItemMainSynchronizeRequest) GetTitle() string {
    return r.title
}

