package ticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票价格库存同步接口 APIRequest
alitrip.ticket.skus.upload

航旅度假新版门票商品价格库存同步接口。
注1、一个票种下可以挂多个规则（规则id必须不一样，每个规则实际对应了一个sku），同一个规则可以在不同票种下使用。
注2、日历库存和区间库存门票，统一使用DateInventory结构。对于日历库存门票请上传每一天的价格库存；对于区间库存门票，建议只上传开始和结束日期的价格库存，也支持上传每天的价格库存，系统会自动进行聚合（取第一天的价格为区间价格，累计所有天的库存为区间库存）。
注3、该接口同时支持 新增某个规则的价格库存 和 更新现有规则的价格库存。如果不清楚是否已在某个规则下上传过价格库存，请使用alitrip.ticket.product.query接口进行查询。如果该规则在该票种下已经存在，则该接口会判断为是价格库存更新操作。
*/
type AlitripTicketSkusUploadRequest struct {
    model.Params

    // 特殊必填，阿里标准收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
    aliProductId   int64 

    // 特殊必填，商户收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
    outProductId   string 

    // 特殊必填，淘宝商品id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
    itemId   int64 

    // 必填，门票 票种类型
    ticketType   string 

    // 可选，门票场次（场次门票专用，对于场次门票必选）
    ticketSeason   string 

    // 可选，门票区域（场次门票专用，对于场次门票必选）
    ticketArea   string 

    // 必填，该票种下使用的价格规则。
    priceRules   []PriceRule 

}

func NewAlitripTicketSkusUploadRequest() *AlitripTicketSkusUploadRequest{
    return &AlitripTicketSkusUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTicketSkusUploadRequest) GetApiMethodName() string {
    return "alitrip.ticket.skus.upload"
}

func (r AlitripTicketSkusUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTicketSkusUploadRequest) SetAliProductId(aliProductId int64) error {
    r.aliProductId = aliProductId
    r.Set("ali_product_id", aliProductId)
    return nil
}

func (r AlitripTicketSkusUploadRequest) GetAliProductId() int64 {
    return r.aliProductId
}

func (r *AlitripTicketSkusUploadRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

func (r AlitripTicketSkusUploadRequest) GetOutProductId() string {
    return r.outProductId
}

func (r *AlitripTicketSkusUploadRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlitripTicketSkusUploadRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlitripTicketSkusUploadRequest) SetTicketType(ticketType string) error {
    r.ticketType = ticketType
    r.Set("ticket_type", ticketType)
    return nil
}

func (r AlitripTicketSkusUploadRequest) GetTicketType() string {
    return r.ticketType
}

func (r *AlitripTicketSkusUploadRequest) SetTicketSeason(ticketSeason string) error {
    r.ticketSeason = ticketSeason
    r.Set("ticket_season", ticketSeason)
    return nil
}

func (r AlitripTicketSkusUploadRequest) GetTicketSeason() string {
    return r.ticketSeason
}

func (r *AlitripTicketSkusUploadRequest) SetTicketArea(ticketArea string) error {
    r.ticketArea = ticketArea
    r.Set("ticket_area", ticketArea)
    return nil
}

func (r AlitripTicketSkusUploadRequest) GetTicketArea() string {
    return r.ticketArea
}

func (r *AlitripTicketSkusUploadRequest) SetPriceRules(priceRules []PriceRule) error {
    r.priceRules = priceRules
    r.Set("price_rules", priceRules)
    return nil
}

func (r AlitripTicketSkusUploadRequest) GetPriceRules() []PriceRule {
    return r.priceRules
}

