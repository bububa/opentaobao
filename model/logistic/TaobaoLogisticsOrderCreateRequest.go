package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建物流订单 API请求
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

// 初始化TaobaoLogisticsOrderCreateRequest对象
func NewTaobaoLogisticsOrderCreateRequest() *TaobaoLogisticsOrderCreateRequest{
    return &TaobaoLogisticsOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOrderCreateRequest) GetApiMethodName() string {
    return "taobao.logistics.order.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LogisType Setter
// 发货方式,默认为自己联系发货。可选值:ONLINE(在线下单)、OFFLINE(自己联系)。
func (r *TaobaoLogisticsOrderCreateRequest) SetLogisType(logisType string) error {
    r.logisType = logisType
    r.Set("logis_type", logisType)
    return nil
}

// LogisType Getter
func (r TaobaoLogisticsOrderCreateRequest) GetLogisType() string {
    return r.logisType
}
// LogisCompanyCode Setter
// 发货的物流公司代码，如申通=STO，圆通=YTO。is_consign=true时，此项必填。
func (r *TaobaoLogisticsOrderCreateRequest) SetLogisCompanyCode(logisCompanyCode string) error {
    r.logisCompanyCode = logisCompanyCode
    r.Set("logis_company_code", logisCompanyCode)
    return nil
}

// LogisCompanyCode Getter
func (r TaobaoLogisticsOrderCreateRequest) GetLogisCompanyCode() string {
    return r.logisCompanyCode
}
// MailNo Setter
// 发货的物流公司运单号。在logis_type=OFFLINE且is_consign=true时，此项必填。
func (r *TaobaoLogisticsOrderCreateRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

// MailNo Getter
func (r TaobaoLogisticsOrderCreateRequest) GetMailNo() string {
    return r.mailNo
}
// GoodsNames Setter
// 运送的货物名称列表，用|号隔开
func (r *TaobaoLogisticsOrderCreateRequest) SetGoodsNames(goodsNames string) error {
    r.goodsNames = goodsNames
    r.Set("goods_names", goodsNames)
    return nil
}

// GoodsNames Getter
func (r TaobaoLogisticsOrderCreateRequest) GetGoodsNames() string {
    return r.goodsNames
}
// GoodsQuantities Setter
// 运送货物的数量列表，用|号隔开
func (r *TaobaoLogisticsOrderCreateRequest) SetGoodsQuantities(goodsQuantities string) error {
    r.goodsQuantities = goodsQuantities
    r.Set("goods_quantities", goodsQuantities)
    return nil
}

// GoodsQuantities Getter
func (r TaobaoLogisticsOrderCreateRequest) GetGoodsQuantities() string {
    return r.goodsQuantities
}
// ItemValues Setter
// 运送货物的单价列表(注意：单位为分），用|号隔开
func (r *TaobaoLogisticsOrderCreateRequest) SetItemValues(itemValues string) error {
    r.itemValues = itemValues
    r.Set("item_values", itemValues)
    return nil
}

// ItemValues Getter
func (r TaobaoLogisticsOrderCreateRequest) GetItemValues() string {
    return r.itemValues
}
// SellerWangwangId Setter
// 卖家旺旺号
func (r *TaobaoLogisticsOrderCreateRequest) SetSellerWangwangId(sellerWangwangId string) error {
    r.sellerWangwangId = sellerWangwangId
    r.Set("seller_wangwang_id", sellerWangwangId)
    return nil
}

// SellerWangwangId Getter
func (r TaobaoLogisticsOrderCreateRequest) GetSellerWangwangId() string {
    return r.sellerWangwangId
}
// TradeId Setter
// 订单的交易号码
func (r *TaobaoLogisticsOrderCreateRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoLogisticsOrderCreateRequest) GetTradeId() int64 {
    return r.tradeId
}
// IsConsign Setter
// 创建订单同时是否进行发货，默认发货。
func (r *TaobaoLogisticsOrderCreateRequest) SetIsConsign(isConsign bool) error {
    r.isConsign = isConsign
    r.Set("is_consign", isConsign)
    return nil
}

// IsConsign Getter
func (r TaobaoLogisticsOrderCreateRequest) GetIsConsign() bool {
    return r.isConsign
}
// Shipping Setter
// 运费承担方式。1为买家承担运费，2为卖家承担运费，其他值为错误参数。
func (r *TaobaoLogisticsOrderCreateRequest) SetShipping(shipping int64) error {
    r.shipping = shipping
    r.Set("shipping", shipping)
    return nil
}

// Shipping Getter
func (r TaobaoLogisticsOrderCreateRequest) GetShipping() int64 {
    return r.shipping
}
