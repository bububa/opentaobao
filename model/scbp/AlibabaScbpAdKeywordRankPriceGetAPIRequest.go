package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordrankpricegetAPIRequest 外贸直通车关键词前五名排价 API请求
// alibaba.scbp.ad.keyword.rank.price.get
//
// 外贸直通车关键词前五名排价
type AlibabascbpadkeywordrankpricegetAPIRequest struct {
	model.Params
	// 关键词
	_keyword string
}

// NewAlibabascbpadkeywordrankpricegetRequest 初始化AlibabascbpadkeywordrankpricegetAPIRequest对象
func NewAlibabascbpadkeywordrankpricegetRequest() *AlibabascbpadkeywordrankpricegetAPIRequest {
	return &AlibabascbpadkeywordrankpricegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordrankpricegetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.rank.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordrankpricegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordrankpricegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 关键词
func (r *AlibabascbpadkeywordrankpricegetAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabascbpadkeywordrankpricegetAPIRequest) GetKeyword() string {
	return r._keyword
}
