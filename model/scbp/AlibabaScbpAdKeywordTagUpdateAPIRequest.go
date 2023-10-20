package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordtagupdateAPIRequest 修改关键词所属分组 API请求
// alibaba.scbp.ad.keyword.tag.update
//
// 修改关键词所属分组
type AlibabascbpadkeywordtagupdateAPIRequest struct {
	model.Params
	// 关键词ID列表
	_keywordIdList []string
	// 关键词分组ID,不传表示取消关键词的分组
	_tagIdList []string
}

// NewAlibabascbpadkeywordtagupdateRequest 初始化AlibabascbpadkeywordtagupdateAPIRequest对象
func NewAlibabascbpadkeywordtagupdateRequest() *AlibabascbpadkeywordtagupdateAPIRequest {
	return &AlibabascbpadkeywordtagupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordtagupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.tag.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordtagupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordtagupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordIdList is KeywordIdList Setter
// 关键词ID列表
func (r *AlibabascbpadkeywordtagupdateAPIRequest) SetKeywordIdList(_keywordIdList []string) error {
	r._keywordIdList = _keywordIdList
	r.Set("keyword_id_list", _keywordIdList)
	return nil
}

// GetKeywordIdList KeywordIdList Getter
func (r AlibabascbpadkeywordtagupdateAPIRequest) GetKeywordIdList() []string {
	return r._keywordIdList
}

// SetTagIdList is TagIdList Setter
// 关键词分组ID,不传表示取消关键词的分组
func (r *AlibabascbpadkeywordtagupdateAPIRequest) SetTagIdList(_tagIdList []string) error {
	r._tagIdList = _tagIdList
	r.Set("tag_id_list", _tagIdList)
	return nil
}

// GetTagIdList TagIdList Getter
func (r AlibabascbpadkeywordtagupdateAPIRequest) GetTagIdList() []string {
	return r._tagIdList
}
