package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordTagUpdateAPIRequest 修改关键词所属分组 API请求
// alibaba.scbp.ad.keyword.tag.update
//
// 修改关键词所属分组
type AlibabaScbpAdKeywordTagUpdateAPIRequest struct {
	model.Params
	// 关键词ID列表
	_keywordIdList []string
	// 关键词分组ID,不传表示取消关键词的分组
	_tagIdList []string
}

// NewAlibabaScbpAdKeywordTagUpdateRequest 初始化AlibabaScbpAdKeywordTagUpdateAPIRequest对象
func NewAlibabaScbpAdKeywordTagUpdateRequest() *AlibabaScbpAdKeywordTagUpdateAPIRequest {
	return &AlibabaScbpAdKeywordTagUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordTagUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.tag.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordTagUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetKeywordIdList is KeywordIdList Setter
// 关键词ID列表
func (r *AlibabaScbpAdKeywordTagUpdateAPIRequest) SetKeywordIdList(_keywordIdList []string) error {
	r._keywordIdList = _keywordIdList
	r.Set("keyword_id_list", _keywordIdList)
	return nil
}

// GetKeywordIdList KeywordIdList Getter
func (r AlibabaScbpAdKeywordTagUpdateAPIRequest) GetKeywordIdList() []string {
	return r._keywordIdList
}

// SetTagIdList is TagIdList Setter
// 关键词分组ID,不传表示取消关键词的分组
func (r *AlibabaScbpAdKeywordTagUpdateAPIRequest) SetTagIdList(_tagIdList []string) error {
	r._tagIdList = _tagIdList
	r.Set("tag_id_list", _tagIdList)
	return nil
}

// GetTagIdList TagIdList Getter
func (r AlibabaScbpAdKeywordTagUpdateAPIRequest) GetTagIdList() []string {
	return r._tagIdList
}
