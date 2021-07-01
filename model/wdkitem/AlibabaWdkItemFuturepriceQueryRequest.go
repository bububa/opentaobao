package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个商品未来价查询接口 API请求
alibaba.wdk.item.futureprice.query

查询单个商品未来价，融合了未来基础售价+未来促销价
*/
type AlibabaWdkItemFuturepriceQueryAPIRequest struct {
    model.Params
    // 渠道店id
    _shopId   int64
    // 商品编码
    _skuCode   string
    // 渠道
    _orderChannelCode   string
    // 开始时间
    _startTime   string
    // 结束时间，结束时间-开始时间不能超过48小时
    _endTime   string
}

// 初始化AlibabaWdkItemFuturepriceQueryAPIRequest对象
func NewAlibabaWdkItemFuturepriceQueryRequest() *AlibabaWdkItemFuturepriceQueryAPIRequest{
    return &AlibabaWdkItemFuturepriceQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemFuturepriceQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.futureprice.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemFuturepriceQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 渠道店id
func (r *AlibabaWdkItemFuturepriceQueryAPIRequest) SetShopId(_shopId int64) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkItemFuturepriceQueryAPIRequest) GetShopId() int64 {
    return r._shopId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemFuturepriceQueryAPIRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemFuturepriceQueryAPIRequest) GetSkuCode() string {
    return r._skuCode
}
// OrderChannelCode Setter
// 渠道
func (r *AlibabaWdkItemFuturepriceQueryAPIRequest) SetOrderChannelCode(_orderChannelCode string) error {
    r._orderChannelCode = _orderChannelCode
    r.Set("order_channel_code", _orderChannelCode)
    return nil
}

// OrderChannelCode Getter
func (r AlibabaWdkItemFuturepriceQueryAPIRequest) GetOrderChannelCode() string {
    return r._orderChannelCode
}
// StartTime Setter
// 开始时间
func (r *AlibabaWdkItemFuturepriceQueryAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r AlibabaWdkItemFuturepriceQueryAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间，结束时间-开始时间不能超过48小时
func (r *AlibabaWdkItemFuturepriceQueryAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r AlibabaWdkItemFuturepriceQueryAPIRequest) GetEndTime() string {
    return r._endTime
}
