package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordbatchdeleteAPIRequest 外贸直通车批量删除关键词 API请求
// alibaba.scbp.ad.keyword.batchdelete
//
// 外贸直通车批量删除关键词
type AlibabascbpadkeywordbatchdeleteAPIRequest struct {
	model.Params
	// 关键词Id列表
	_keywordIdList []string
}

// NewAlibabascbpadkeywordbatchdeleteRequest 初始化AlibabascbpadkeywordbatchdeleteAPIRequest对象
func NewAlibabascbpadkeywordbatchdeleteRequest() *AlibabascbpadkeywordbatchdeleteAPIRequest {
	return &AlibabascbpadkeywordbatchdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordbatchdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.batchdelete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordbatchdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordbatchdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordIdList is KeywordIdList Setter
// 关键词Id列表
func (r *AlibabascbpadkeywordbatchdeleteAPIRequest) SetKeywordIdList(_keywordIdList []string) error {
	r._keywordIdList = _keywordIdList
	r.Set("keyword_id_list", _keywordIdList)
	return nil
}

// GetKeywordIdList KeywordIdList Getter
func (r AlibabascbpadkeywordbatchdeleteAPIRequest) GetKeywordIdList() []string {
	return r._keywordIdList
}
