package ticket

import (
	"net/url"
	"sync"

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
	// 代表业务来源，gongxiao代表供销平台业务
	_pageSource string
	// 阿里标准收费项目id。与out_product_id，item_id 三者至少填写一个
	_aliProductId int64
	// 商品ID。与out_product_id，ali_product_id三者至少填写一个
	_itemId int64
}

// NewAlitripTicketProductQueryRequest 初始化AlitripTicketProductQueryAPIRequest对象
func NewAlitripTicketProductQueryRequest() *AlitripTicketProductQueryAPIRequest {
	return &AlitripTicketProductQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTicketProductQueryAPIRequest) Reset() {
	r._outProductId = ""
	r._pageSource = ""
	r._aliProductId = 0
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTicketProductQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.ticket.product.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTicketProductQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTicketProductQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutProductId is OutProductId Setter
// 商户自定义收费项目编码。与ali_product_id，item_id 三者至少填写一个
func (r *AlitripTicketProductQueryAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r AlitripTicketProductQueryAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetPageSource is PageSource Setter
// 代表业务来源，gongxiao代表供销平台业务
func (r *AlitripTicketProductQueryAPIRequest) SetPageSource(_pageSource string) error {
	r._pageSource = _pageSource
	r.Set("page_source", _pageSource)
	return nil
}

// GetPageSource PageSource Getter
func (r AlitripTicketProductQueryAPIRequest) GetPageSource() string {
	return r._pageSource
}

// SetAliProductId is AliProductId Setter
// 阿里标准收费项目id。与out_product_id，item_id 三者至少填写一个
func (r *AlitripTicketProductQueryAPIRequest) SetAliProductId(_aliProductId int64) error {
	r._aliProductId = _aliProductId
	r.Set("ali_product_id", _aliProductId)
	return nil
}

// GetAliProductId AliProductId Getter
func (r AlitripTicketProductQueryAPIRequest) GetAliProductId() int64 {
	return r._aliProductId
}

// SetItemId is ItemId Setter
// 商品ID。与out_product_id，ali_product_id三者至少填写一个
func (r *AlitripTicketProductQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripTicketProductQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolAlitripTicketProductQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTicketProductQueryRequest()
	},
}

// GetAlitripTicketProductQueryRequest 从 sync.Pool 获取 AlitripTicketProductQueryAPIRequest
func GetAlitripTicketProductQueryAPIRequest() *AlitripTicketProductQueryAPIRequest {
	return poolAlitripTicketProductQueryAPIRequest.Get().(*AlitripTicketProductQueryAPIRequest)
}

// ReleaseAlitripTicketProductQueryAPIRequest 将 AlitripTicketProductQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripTicketProductQueryAPIRequest(v *AlitripTicketProductQueryAPIRequest) {
	v.Reset()
	poolAlitripTicketProductQueryAPIRequest.Put(v)
}
