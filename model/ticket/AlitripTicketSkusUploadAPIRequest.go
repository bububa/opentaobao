package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketSkusUploadAPIRequest 【门票API2.0】门票价格库存同步接口 API请求
// alitrip.ticket.skus.upload
//
// 航旅度假新版门票商品价格库存同步接口。
// 注1、一个票种下可以挂多个规则（规则id必须不一样，每个规则实际对应了一个sku），同一个规则可以在不同票种下使用。
// 注2、日历库存和区间库存门票，统一使用DateInventory结构。对于日历库存门票请上传每一天的价格库存；对于区间库存门票，建议只上传开始和结束日期的价格库存，也支持上传每天的价格库存，系统会自动进行聚合（取第一天的价格为区间价格，累计所有天的库存为区间库存）。
// 注3、该接口同时支持 新增某个规则的价格库存 和 更新现有规则的价格库存。如果不清楚是否已在某个规则下上传过价格库存，请使用alitrip.ticket.product.query接口进行查询。如果该规则在该票种下已经存在，则该接口会判断为是价格库存更新操作。
type AlitripTicketSkusUploadAPIRequest struct {
	model.Params
	// 特殊必填，阿里标准收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
	_aliProductId int64
	// 特殊必填，商户收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
	_outProductId string
	// 特殊必填，淘宝商品id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
	_itemId int64
	// 必填，门票 票种类型
	_ticketType string
	// 可选，门票场次（场次门票专用，对于场次门票必选）
	_ticketSeason string
	// 可选，门票区域（场次门票专用，对于场次门票必选）
	_ticketArea string
	// 必填，该票种下使用的价格规则。
	_priceRules []PriceRule
}

// NewAlitripTicketSkusUploadRequest 初始化AlitripTicketSkusUploadAPIRequest对象
func NewAlitripTicketSkusUploadRequest() *AlitripTicketSkusUploadAPIRequest {
	return &AlitripTicketSkusUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTicketSkusUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.ticket.skus.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTicketSkusUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAliProductId is AliProductId Setter
// 特殊必填，阿里标准收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
func (r *AlitripTicketSkusUploadAPIRequest) SetAliProductId(_aliProductId int64) error {
	r._aliProductId = _aliProductId
	r.Set("ali_product_id", _aliProductId)
	return nil
}

// GetAliProductId AliProductId Getter
func (r AlitripTicketSkusUploadAPIRequest) GetAliProductId() int64 {
	return r._aliProductId
}

// SetOutProductId is OutProductId Setter
// 特殊必填，商户收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
func (r *AlitripTicketSkusUploadAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r AlitripTicketSkusUploadAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 特殊必填，淘宝商品id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
func (r *AlitripTicketSkusUploadAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripTicketSkusUploadAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetTicketType is TicketType Setter
// 必填，门票 票种类型
func (r *AlitripTicketSkusUploadAPIRequest) SetTicketType(_ticketType string) error {
	r._ticketType = _ticketType
	r.Set("ticket_type", _ticketType)
	return nil
}

// GetTicketType TicketType Getter
func (r AlitripTicketSkusUploadAPIRequest) GetTicketType() string {
	return r._ticketType
}

// SetTicketSeason is TicketSeason Setter
// 可选，门票场次（场次门票专用，对于场次门票必选）
func (r *AlitripTicketSkusUploadAPIRequest) SetTicketSeason(_ticketSeason string) error {
	r._ticketSeason = _ticketSeason
	r.Set("ticket_season", _ticketSeason)
	return nil
}

// GetTicketSeason TicketSeason Getter
func (r AlitripTicketSkusUploadAPIRequest) GetTicketSeason() string {
	return r._ticketSeason
}

// SetTicketArea is TicketArea Setter
// 可选，门票区域（场次门票专用，对于场次门票必选）
func (r *AlitripTicketSkusUploadAPIRequest) SetTicketArea(_ticketArea string) error {
	r._ticketArea = _ticketArea
	r.Set("ticket_area", _ticketArea)
	return nil
}

// GetTicketArea TicketArea Getter
func (r AlitripTicketSkusUploadAPIRequest) GetTicketArea() string {
	return r._ticketArea
}

// SetPriceRules is PriceRules Setter
// 必填，该票种下使用的价格规则。
func (r *AlitripTicketSkusUploadAPIRequest) SetPriceRules(_priceRules []PriceRule) error {
	r._priceRules = _priceRules
	r.Set("price_rules", _priceRules)
	return nil
}

// GetPriceRules PriceRules Getter
func (r AlitripTicketSkusUploadAPIRequest) GetPriceRules() []PriceRule {
	return r._priceRules
}
