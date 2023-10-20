package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordrankgetAPIRequest 获取外贸直通车关键词预估排名 API请求
// alibaba.scbp.ad.keyword.rank.get
//
// 获取外贸直通车关键词预估排名
type AlibabascbpadkeywordrankgetAPIRequest struct {
	model.Params
	// 查询预估排名的关键词
	_keyword string
}

// NewAlibabascbpadkeywordrankgetRequest 初始化AlibabascbpadkeywordrankgetAPIRequest对象
func NewAlibabascbpadkeywordrankgetRequest() *AlibabascbpadkeywordrankgetAPIRequest {
	return &AlibabascbpadkeywordrankgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordrankgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.rank.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordrankgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordrankgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 查询预估排名的关键词
func (r *AlibabascbpadkeywordrankgetAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabascbpadkeywordrankgetAPIRequest) GetKeyword() string {
	return r._keyword
}
