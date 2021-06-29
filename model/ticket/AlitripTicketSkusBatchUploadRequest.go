package ticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票价格库存同步接口（多票种批量更新） API请求
alitrip.ticket.skus.batch.upload

飞猪度假新版门票商品价格库存同步接口（多票种批量更新）。
注1、一个票种下可以挂多个规则（规则id必须不一样，每个规则实际对应了一个sku），同一个规则可以在不同票种下使用。
注2、日历库存和区间库存门票，统一使用DateInventory结构。对于日历库存门票请上传每一天的价格库存；对于区间库存门票，建议只上传开始和结束日期的价格库存，也支持上传每天的价格库存，系统会自动进行聚合（取第一天的价格为区间价格，累计所有天的库存为区间库存）。
注3、该接口同时支持 新增某个规则的价格库存 和 更新现有规则的价格库存。如果不清楚是否已在某个规则下上传过价格库存，请使用alitrip.ticket.product.query接口进行查询。如果该规则在该票种下已经存在，则该接口会判断为是价格库存更新操作。
*/
type AlitripTicketSkusBatchUploadRequest struct {
    model.Params
    // 特殊必填，阿里标准收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
    aliProductId   int64
    // 特殊必填，商户收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
    outProductId   string
    // 特殊必填，淘宝商品id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
    itemId   int64
    // 必填，各票种下sku的价格库存参数。
    ticketPriceRules   []TicketPriceRule
}

// 初始化AlitripTicketSkusBatchUploadRequest对象
func NewAlitripTicketSkusBatchUploadRequest() *AlitripTicketSkusBatchUploadRequest{
    return &AlitripTicketSkusBatchUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTicketSkusBatchUploadRequest) GetApiMethodName() string {
    return "alitrip.ticket.skus.batch.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTicketSkusBatchUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AliProductId Setter
// 特殊必填，阿里标准收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
func (r *AlitripTicketSkusBatchUploadRequest) SetAliProductId(aliProductId int64) error {
    r.aliProductId = aliProductId
    r.Set("ali_product_id", aliProductId)
    return nil
}

// AliProductId Getter
func (r AlitripTicketSkusBatchUploadRequest) GetAliProductId() int64 {
    return r.aliProductId
}
// OutProductId Setter
// 特殊必填，商户收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
func (r *AlitripTicketSkusBatchUploadRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

// OutProductId Getter
func (r AlitripTicketSkusBatchUploadRequest) GetOutProductId() string {
    return r.outProductId
}
// ItemId Setter
// 特殊必填，淘宝商品id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
func (r *AlitripTicketSkusBatchUploadRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlitripTicketSkusBatchUploadRequest) GetItemId() int64 {
    return r.itemId
}
// TicketPriceRules Setter
// 必填，各票种下sku的价格库存参数。
func (r *AlitripTicketSkusBatchUploadRequest) SetTicketPriceRules(ticketPriceRules []TicketPriceRule) error {
    r.ticketPriceRules = ticketPriceRules
    r.Set("ticket_price_rules", ticketPriceRules)
    return nil
}

// TicketPriceRules Getter
func (r AlitripTicketSkusBatchUploadRequest) GetTicketPriceRules() []TicketPriceRule {
    return r.ticketPriceRules
}
