package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSearchAPIRequest 根据发票抬头搜索发票 API请求
// alitrip.btrip.invoice.search
//
// 用户根据发票抬头搜索发票信息
type AlitripBtripInvoiceSearchAPIRequest struct {
	model.Params
	// 企业id
	_corpId string
	// 用户id
	_userId string
	// 发票抬头
	_title string
}

// NewAlitripBtripInvoiceSearchRequest 初始化AlitripBtripInvoiceSearchAPIRequest对象
func NewAlitripBtripInvoiceSearchRequest() *AlitripBtripInvoiceSearchAPIRequest {
	return &AlitripBtripInvoiceSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CorpId Setter
// 企业id
func (r *AlitripBtripInvoiceSearchAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// Get CorpId Getter
func (r AlitripBtripInvoiceSearchAPIRequest) GetCorpId() string {
	return r._corpId
}

// Set is UserId Setter
// 用户id
func (r *AlitripBtripInvoiceSearchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlitripBtripInvoiceSearchAPIRequest) GetUserId() string {
	return r._userId
}

// Set is Title Setter
// 发票抬头
func (r *AlitripBtripInvoiceSearchAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r AlitripBtripInvoiceSearchAPIRequest) GetTitle() string {
	return r._title
}
