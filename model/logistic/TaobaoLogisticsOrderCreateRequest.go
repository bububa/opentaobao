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
    _logisType   string
    // 发货的物流公司代码，如申通=STO，圆通=YTO。is_consign=true时，此项必填。
    _logisCompanyCode   string
    // 发货的物流公司运单号。在logis_type=OFFLINE且is_consign=true时，此项必填。
    _mailNo   string
    // 运送的货物名称列表，用|号隔开
    _goodsNames   string
    // 运送货物的数量列表，用|号隔开
    _goodsQuantities   string
    // 运送货物的单价列表(注意：单位为分），用|号隔开
    _itemValues   string
    // 卖家旺旺号
    _sellerWangwangId   string
    // 订单的交易号码
    _tradeId   int64
    // 创建订单同时是否进行发货，默认发货。
    _isConsign   bool
    // 运费承担方式。1为买家承担运费，2为卖家承担运费，其他值为错误参数。
    _shipping   int64
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
func (r *TaobaoLogisticsOrderCreateRequest) SetLogisType(_logisType string) error {
    r._logisType = _logisType
    r.Set("logis_type", _logisType)
    return nil
}

// LogisType Getter
func (r TaobaoLogisticsOrderCreateRequest) GetLogisType() string {
    return r._logisType
}
// LogisCompanyCode Setter
// 发货的物流公司代码，如申通=STO，圆通=YTO。is_consign=true时，此项必填。
func (r *TaobaoLogisticsOrderCreateRequest) SetLogisCompanyCode(_logisCompanyCode string) error {
    r._logisCompanyCode = _logisCompanyCode
    r.Set("logis_company_code", _logisCompanyCode)
    return nil
}

// LogisCompanyCode Getter
func (r TaobaoLogisticsOrderCreateRequest) GetLogisCompanyCode() string {
    return r._logisCompanyCode
}
// MailNo Setter
// 发货的物流公司运单号。在logis_type=OFFLINE且is_consign=true时，此项必填。
func (r *TaobaoLogisticsOrderCreateRequest) SetMailNo(_mailNo string) error {
    r._mailNo = _mailNo
    r.Set("mail_no", _mailNo)
    return nil
}

// MailNo Getter
func (r TaobaoLogisticsOrderCreateRequest) GetMailNo() string {
    return r._mailNo
}
// GoodsNames Setter
// 运送的货物名称列表，用|号隔开
func (r *TaobaoLogisticsOrderCreateRequest) SetGoodsNames(_goodsNames string) error {
    r._goodsNames = _goodsNames
    r.Set("goods_names", _goodsNames)
    return nil
}

// GoodsNames Getter
func (r TaobaoLogisticsOrderCreateRequest) GetGoodsNames() string {
    return r._goodsNames
}
// GoodsQuantities Setter
// 运送货物的数量列表，用|号隔开
func (r *TaobaoLogisticsOrderCreateRequest) SetGoodsQuantities(_goodsQuantities string) error {
    r._goodsQuantities = _goodsQuantities
    r.Set("goods_quantities", _goodsQuantities)
    return nil
}

// GoodsQuantities Getter
func (r TaobaoLogisticsOrderCreateRequest) GetGoodsQuantities() string {
    return r._goodsQuantities
}
// ItemValues Setter
// 运送货物的单价列表(注意：单位为分），用|号隔开
func (r *TaobaoLogisticsOrderCreateRequest) SetItemValues(_itemValues string) error {
    r._itemValues = _itemValues
    r.Set("item_values", _itemValues)
    return nil
}

// ItemValues Getter
func (r TaobaoLogisticsOrderCreateRequest) GetItemValues() string {
    return r._itemValues
}
// SellerWangwangId Setter
// 卖家旺旺号
func (r *TaobaoLogisticsOrderCreateRequest) SetSellerWangwangId(_sellerWangwangId string) error {
    r._sellerWangwangId = _sellerWangwangId
    r.Set("seller_wangwang_id", _sellerWangwangId)
    return nil
}

// SellerWangwangId Getter
func (r TaobaoLogisticsOrderCreateRequest) GetSellerWangwangId() string {
    return r._sellerWangwangId
}
// TradeId Setter
// 订单的交易号码
func (r *TaobaoLogisticsOrderCreateRequest) SetTradeId(_tradeId int64) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoLogisticsOrderCreateRequest) GetTradeId() int64 {
    return r._tradeId
}
// IsConsign Setter
// 创建订单同时是否进行发货，默认发货。
func (r *TaobaoLogisticsOrderCreateRequest) SetIsConsign(_isConsign bool) error {
    r._isConsign = _isConsign
    r.Set("is_consign", _isConsign)
    return nil
}

// IsConsign Getter
func (r TaobaoLogisticsOrderCreateRequest) GetIsConsign() bool {
    return r._isConsign
}
// Shipping Setter
// 运费承担方式。1为买家承担运费，2为卖家承担运费，其他值为错误参数。
func (r *TaobaoLogisticsOrderCreateRequest) SetShipping(_shipping int64) error {
    r._shipping = _shipping
    r.Set("shipping", _shipping)
    return nil
}

// Shipping Getter
func (r TaobaoLogisticsOrderCreateRequest) GetShipping() int64 {
    return r._shipping
}
