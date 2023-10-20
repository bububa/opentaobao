package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripticketproductqueryAPIRequest 【门票API2.0】门票商品查询接口 API请求
// alitrip.ticket.product.query
//
// 门票商品查询接口：返回商家上传的门票商品信息
type AlitripticketproductqueryAPIRequest struct {
	model.Params
	// 商户自定义收费项目编码。与ali_product_id，item_id 三者至少填写一个
	_outProductId string
	// 代表业务来源，gongxiao代表供销平台业务
	_pageSource string
	// 阿里标准收费项目id。与out_product_id，item_id 三者至少填写一个
	_aliProductId int64
	// 商品ID。与out_product_id，ali_product_id三者至少填写一个
	_itemId int64
}

// NewAlitripticketproductqueryRequest 初始化AlitripticketproductqueryAPIRequest对象
func NewAlitripticketproductqueryRequest() *AlitripticketproductqueryAPIRequest {
	return &AlitripticketproductqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripticketproductqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.ticket.product.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripticketproductqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripticketproductqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutProductId is OutProductId Setter
// 商户自定义收费项目编码。与ali_product_id，item_id 三者至少填写一个
func (r *AlitripticketproductqueryAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r AlitripticketproductqueryAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetPageSource is PageSource Setter
// 代表业务来源，gongxiao代表供销平台业务
func (r *AlitripticketproductqueryAPIRequest) SetPageSource(_pageSource string) error {
	r._pageSource = _pageSource
	r.Set("page_source", _pageSource)
	return nil
}

// GetPageSource PageSource Getter
func (r AlitripticketproductqueryAPIRequest) GetPageSource() string {
	return r._pageSource
}

// SetAliProductId is AliProductId Setter
// 阿里标准收费项目id。与out_product_id，item_id 三者至少填写一个
func (r *AlitripticketproductqueryAPIRequest) SetAliProductId(_aliProductId int64) error {
	r._aliProductId = _aliProductId
	r.Set("ali_product_id", _aliProductId)
	return nil
}

// GetAliProductId AliProductId Getter
func (r AlitripticketproductqueryAPIRequest) GetAliProductId() int64 {
	return r._aliProductId
}

// SetItemId is ItemId Setter
// 商品ID。与out_product_id，ali_product_id三者至少填写一个
func (r *AlitripticketproductqueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripticketproductqueryAPIRequest) GetItemId() int64 {
	return r._itemId
}
