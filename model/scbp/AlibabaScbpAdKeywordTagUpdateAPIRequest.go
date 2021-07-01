package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordTagUpdateAPIRequest
修改关键词所属分组 API请求
alibaba.scbp.ad.keyword.tag.update

修改关键词所属分组 */
type AlibabaScbpAdKeywordTagUpdateAPIRequest struct {
	model.Params
	// 关键词ID列表
	_keywordIdList []int64
	// 关键词分组ID,不传表示取消关键词的分组
	_tagIdList []int64
}

// New
