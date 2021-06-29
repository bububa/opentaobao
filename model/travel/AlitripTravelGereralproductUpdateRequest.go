package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通用类目产品发布编辑 API请求
alitrip.travel.gereralproduct.update

提供给飞猪供销平台供应商发布编辑通用类目产品的API
*/
type AlitripTravelGereralproductUpdateRequest struct {
    model.Params
    // 产品基本信息
    _baseInfo   *GeneralProductBaseInfo
    // 退款规则结构
    _refundInfo   *ItemRefundInfo
    // 必填，预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
    _bookingRules   []BookingRuleInfo
    // 产品销售信息
    _productSaleInfo   *ProductSaleInfo
    // 更新sku信息，仅限日历商品使用
    _dateSkuInfoList   []DateSkuInfo
}

// 初始化AlitripTravelGereralproductUpdateRequest对象
func NewAlitripTravelGereralproductUpdateRequest() *AlitripTravelGereralproductUpdateRequest{
    return &AlitripTravelGereralproductUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelGereralproductUpdateRequest) GetApiMethodName() string {
    return "alitrip.travel.gereralproduct.update"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelGereralproductUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseInfo Setter
// 产品基本信息
func (r *AlitripTravelGereralproductUpdateRequest) SetBaseInfo(_baseInfo *GeneralProductBaseInfo) error {
    r._baseInfo = _baseInfo
    r.Set("base_info", _baseInfo)
    return nil
}

// BaseInfo Getter
func (r AlitripTravelGereralproductUpdateRequest) GetBaseInfo() *GeneralProductBaseInfo {
    return r._baseInfo
}
// RefundInfo Setter
// 退款规则结构
func (r *AlitripTravelGereralproductUpdateRequest) SetRefundInfo(_refundInfo *ItemRefundInfo) error {
    r._refundInfo = _refundInfo
    r.Set("refund_info", _refundInfo)
    return nil
}

// RefundInfo Getter
func (r AlitripTravelGereralproductUpdateRequest) GetRefundInfo() *ItemRefundInfo {
    return r._refundInfo
}
// BookingRules Setter
// 必填，预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
func (r *AlitripTravelGereralproductUpdateRequest) SetBookingRules(_bookingRules []BookingRuleInfo) error {
    r._bookingRules = _bookingRules
    r.Set("booking_rules", _bookingRules)
    return nil
}

// BookingRules Getter
func (r AlitripTravelGereralproductUpdateRequest) GetBookingRules() []BookingRuleInfo {
    return r._bookingRules
}
// ProductSaleInfo Setter
// 产品销售信息
func (r *AlitripTravelGereralproductUpdateRequest) SetProductSaleInfo(_productSaleInfo *ProductSaleInfo) error {
    r._productSaleInfo = _productSaleInfo
    r.Set("product_sale_info", _productSaleInfo)
    return nil
}

// ProductSaleInfo Getter
func (r AlitripTravelGereralproductUpdateRequest) GetProductSaleInfo() *ProductSaleInfo {
    return r._productSaleInfo
}
// DateSkuInfoList Setter
// 更新sku信息，仅限日历商品使用
func (r *AlitripTravelGereralproductUpdateRequest) SetDateSkuInfoList(_dateSkuInfoList []DateSkuInfo) error {
    r._dateSkuInfoList = _dateSkuInfoList
    r.Set("date_sku_info_list", _dateSkuInfoList)
    return nil
}

// DateSkuInfoList Getter
func (r AlitripTravelGereralproductUpdateRequest) GetDateSkuInfoList() []DateSkuInfo {
    return r._dateSkuInfoList
}
