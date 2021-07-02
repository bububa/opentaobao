package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketProductQueryAPIRequest 【门票API2.0】门票商品查询接口 API请求
// alitrip.ticket.product.query
//
// 门票商品查询接口：返回商家上传的门票商品信息
type AlitripTicketProductQueryAPIRequest struct {
	model.Params
	// 商户自定义收费项目编码。与ali_product_id，item_id 三者至少填写一个
	_outProductId string
	// 阿里标准收费项目id。与out_product_id，item_id 三者至少填写一个
	_aliProductId int64
	// 商品ID。与out_product_id，ali_product_id三者至少填写一个
	_itemId int64
	// 代表业务来源，gongxiao代表供销平台业务
	_pageSource string
}

// NewAlitripTicketProductQueryRequest 初始化AlitripTicketProductQueryAPIRequest对象
func NewAlitripTicketProductQueryRequest() *AlitripTicketProductQueryAPIRequest {
	return &AlitripTicketProductQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTicketProductQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.ticket.product.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTicketProductQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OutProductId Setter
// 商户自定义收费项目编码。与ali_product_id，item_id 三者至少填写一个
func (r *AlitripTicketProductQueryAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// Get OutProductId Getter
func (r AlitripTicketProductQueryAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// Set is AliProductId Setter
// 阿里标准收费项目id。与out_product_id，item_id 三者至少填写一个
func (r *AlitripTicketProductQueryAPIRequest) SetAliProductId(_aliProductId int64) error {
	r._aliProductId = _aliProductId
	r.Set("ali_product_id", _aliProductId)
	return nil
}

// Get AliProductId Getter
func (r AlitripTicketProductQueryAPIRequest) GetAliProductId() int64 {
	return r._aliProductId
}

// Set is ItemId Setter
// 商品ID。与out_product_id，ali_product_id三者至少填写一个
func (r *AlitripTicketProductQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r AlitripTicketProductQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is PageSource Setter
// 代表业务来源，gongxiao代表供销平台业务
func (r *AlitripTicketProductQueryAPIRequest) SetPageSource(_pageSource string) error {
	r._pageSource = _pageSource
	r.Set("page_source", _pageSource)
	return nil
}

// Get PageSource Getter
func (r AlitripTicketProductQueryAPIRequest) GetPageSource() string {
	return r._pageSource
}
