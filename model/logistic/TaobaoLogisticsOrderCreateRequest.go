package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建物流订单 APIRequest
taobao.logistics.order.create

用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。
*/
type TaobaoLogisticsOrderCreateRequest struct {
    model.Params

    // 发货方式,默认为自己联系发货。可选值:ONLINE(在线下单)、OFFLINE(自己联系)。
    logisType   string 

    // 发货的物流公司代码，如申通=STO，圆通=YTO。is_consign=true时，此项必填。
    logisCompanyCode   string 

    // 发货的物流公司运单号。在logis_type=OFFLINE且is_consign=true时，此项必填。
    mailNo   string 

    // 运送的货物名称列表，用|号隔开
    goodsNames   string 

    // 运送货物的数量列表，用|号隔开
    goodsQuantities   string 

    // 运送货物的单价列表(注意：单位为分），用|号隔开
    itemValues   string 

    // 卖家旺旺号
    sellerWangwangId   string 

    // 订单的交易号码
    tradeId   int64 

    // 创建订单同时是否进行发货，默认发货。
    isConsign   bool 

    // 运费承担方式。1为买家承担运费，2为卖家承担运费，其他值为错误参数。
    shipping   int64 

}

func NewTaobaoLogisticsOrderCreateRequest() *TaobaoLogisticsOrderCreateRequest{
    return &TaobaoLogisticsOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsOrderCreateRequest) GetApiMethodName() string {
    return "taobao.logistics.order.create"
}

func (r TaobaoLogisticsOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsOrderCreateRequest) SetLogisType(logisType string) error {
    r.logisType = logisType
    r.Set("logis_type", logisType)
    return nil
}

func (r TaobaoLogisticsOrderCreateRequest) GetLogisType() string {
    return r.logisType
}

func (r *TaobaoLogisticsOrderCreateRequest) SetLogisCompanyCode(logisCompanyCode string) error {
    r.logisCompanyCode = logisCompanyCode
    r.Set("logis_company_code", logisCompanyCode)
    return nil
}

func (r TaobaoLogisticsOrderCreateRequest) GetLogisCompanyCode() string {
    return r.logisCompanyCode
}

func (r *TaobaoLogisticsOrderCreateRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

func (r TaobaoLogisticsOrderCreateRequest) GetMailNo() string {
    return r.mailNo
}

func (r *TaobaoLogisticsOrderCreateRequest) SetGoodsNames(goodsNames string) error {
    r.goodsNames = goodsNames
    r.Set("goods_names", goodsNames)
    return nil
}

func (r TaobaoLogisticsOrderCreateRequest) GetGoodsNames() string {
    return r.goodsNames
}

func (r *TaobaoLogisticsOrderCreateRequest) SetGoodsQuantities(goodsQuantities string) error {
    r.goodsQuantities = goodsQuantities
    r.Set("goods_quantities", goodsQuantities)
    return nil
}

func (r TaobaoLogisticsOrderCreateRequest) GetGoodsQuantities() string {
    return r.goodsQuantities
}

func (r *TaobaoLogisticsOrderCreateRequest) SetItemValues(itemValues string) error {
    r.itemValues = itemValues
    r.Set("item_values", itemValues)
    return nil
}

func (r TaobaoLogisticsOrderCreateRequest) GetItemValues() string {
    return r.itemValues
}

func (r *TaobaoLogisticsOrderCreateRequest) SetSellerWangwangId(sellerWangwangId string) error {
    r.sellerWangwangId = sellerWangwangId
    r.Set("seller_wangwang_id", sellerWangwangId)
    return nil
}

func (r TaobaoLogisticsOrderCreateRequest) GetSellerWangwangId() string {
    return r.sellerWangwangId
}

func (r *TaobaoLogisticsOrderCreateRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r TaobaoLogisticsOrderCreateRequest) GetTradeId() int64 {
    return r.tradeId
}

func (r *TaobaoLogisticsOrderCreateRequest) SetIsConsign(isConsign bool) error {
    r.isConsign = isConsign
    r.Set("is_consign", isConsign)
    return nil
}

func (r TaobaoLogisticsOrderCreateRequest) GetIsConsign() bool {
    return r.isConsign
}

func (r *TaobaoLogisticsOrderCreateRequest) SetShipping(shipping int64) error {
    r.shipping = shipping
    r.Set("shipping", shipping)
    return nil
}

func (r TaobaoLogisticsOrderCreateRequest) GetShipping() int64 {
    return r.shipping
}

