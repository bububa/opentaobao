package ticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketSkusBatchUploadAPIRequest 【门票API2.0】门票价格库存同步接口（多票种批量更新） API请求
// alitrip.ticket.skus.batch.upload
//
// 飞猪度假新版门票商品价格库存同步接口（多票种批量更新）。
// 注1、一个票种下可以挂多个规则（规则id必须不一样，每个规则实际对应了一个sku），同一个规则可以在不同票种下使用。
// 注2、日历库存和区间库存门票，统一使用DateInventory结构。对于日历库存门票请上传每一天的价格库存；对于区间库存门票，建议只上传开始和结束日期的价格库存，也支持上传每天的价格库存，系统会自动进行聚合（取第一天的价格为区间价格，累计所有天的库存为区间库存）。
// 注3、该接口同时支持 新增某个规则的价格库存 和 更新现有规则的价格库存。如果不清楚是否已在某个规则下上传过价格库存，请使用alitrip.ticket.product.query接口进行查询。如果该规则在该票种下已经存在，则该接口会判断为是价格库存更新操作。
type AlitripTicketSkusBatchUploadAPIRequest struct {
	model.Params
	// 必填，各票种下sku的价格库存参数。
	_ticketPriceRules []TicketPriceRule
	// 特殊必填，商户收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
	_outProductId string
	// 特殊必填，阿里标准收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
	_aliProductId int64
	// 特殊必填，淘宝商品id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
	_itemId int64
}

// NewAlitripTicketSkusBatchUploadRequest 初始化AlitripTicketSkusBatchUploadAPIRequest对象
func NewAlitripTicketSkusBatchUploadRequest() *AlitripTicketSkusBatchUploadAPIRequest {
	return &AlitripTicketSkusBatchUploadAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTicketSkusBatchUploadAPIRequest) Reset() {
	r._ticketPriceRules = r._ticketPriceRules[:0]
	r._outProductId = ""
	r._aliProductId = 0
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTicketSkusBatchUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.ticket.skus.batch.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTicketSkusBatchUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTicketSkusBatchUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketPriceRules is TicketPriceRules Setter
// 必填，各票种下sku的价格库存参数。
func (r *AlitripTicketSkusBatchUploadAPIRequest) SetTicketPriceRules(_ticketPriceRules []TicketPriceRule) error {
	r._ticketPriceRules = _ticketPriceRules
	r.Set("ticket_price_rules", _ticketPriceRules)
	return nil
}

// GetTicketPriceRules TicketPriceRules Getter
func (r AlitripTicketSkusBatchUploadAPIRequest) GetTicketPriceRules() []TicketPriceRule {
	return r._ticketPriceRules
}

// SetOutProductId is OutProductId Setter
// 特殊必填，商户收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
func (r *AlitripTicketSkusBatchUploadAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r AlitripTicketSkusBatchUploadAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetAliProductId is AliProductId Setter
// 特殊必填，阿里标准收费项目id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
func (r *AlitripTicketSkusBatchUploadAPIRequest) SetAliProductId(_aliProductId int64) error {
	r._aliProductId = _aliProductId
	r.Set("ali_product_id", _aliProductId)
	return nil
}

// GetAliProductId AliProductId Getter
func (r AlitripTicketSkusBatchUploadAPIRequest) GetAliProductId() int64 {
	return r._aliProductId
}

// SetItemId is ItemId Setter
// 特殊必填，淘宝商品id。ali_product_id, item_id与out_product_id三选一，至少填写其中一个
func (r *AlitripTicketSkusBatchUploadAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripTicketSkusBatchUploadAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolAlitripTicketSkusBatchUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTicketSkusBatchUploadRequest()
	},
}

// GetAlitripTicketSkusBatchUploadRequest 从 sync.Pool 获取 AlitripTicketSkusBatchUploadAPIRequest
func GetAlitripTicketSkusBatchUploadAPIRequest() *AlitripTicketSkusBatchUploadAPIRequest {
	return poolAlitripTicketSkusBatchUploadAPIRequest.Get().(*AlitripTicketSkusBatchUploadAPIRequest)
}

// ReleaseAlitripTicketSkusBatchUploadAPIRequest 将 AlitripTicketSkusBatchUploadAPIRequest 放入 sync.Pool
func ReleaseAlitripTicketSkusBatchUploadAPIRequest(v *AlitripTicketSkusBatchUploadAPIRequest) {
	v.Reset()
	poolAlitripTicketSkusBatchUploadAPIRequest.Put(v)
}
