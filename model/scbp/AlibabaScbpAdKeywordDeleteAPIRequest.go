package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeyworddeleteAPIRequest 外贸直通车删除关键词 API请求
// alibaba.scbp.ad.keyword.delete
//
// 外贸直通车删除关键词
type AlibabascbpadkeyworddeleteAPIRequest struct {
	model.Params
	// 要删除的关键词
	_adKeyword string
}

// NewAlibabascbpadkeyworddeleteRequest 初始化AlibabascbpadkeyworddeleteAPIRequest对象
func NewAlibabascbpadkeyworddeleteRequest() *AlibabascbpadkeyworddeleteAPIRequest {
	return &AlibabascbpadkeyworddeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeyworddeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeyworddeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeyworddeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdKeyword is AdKeyword Setter
// 要删除的关键词
func (r *AlibabascbpadkeyworddeleteAPIRequest) SetAdKeyword(_adKeyword string) error {
	r._adKeyword = _adKeyword
	r.Set("ad_keyword", _adKeyword)
	return nil
}

// GetAdKeyword AdKeyword Getter
func (r AlibabascbpadkeyworddeleteAPIRequest) GetAdKeyword() string {
	return r._adKeyword
}
