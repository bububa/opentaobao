package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordidsDeletedGetAPIResponse 获取删除的词ID API返回值
// taobao.simba.keywordids.deleted.get
//
// 获取删除的词ID
type TaobaoSimbaKeywordidsDeletedGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordidsDeletedGetAPIResponseModel
}

// TaobaoSimbaKeywordidsDeletedGetAPIResponseModel is 获取删除的词ID 成功返回结果
type TaobaoSimbaKeywordidsDeletedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywordids_deleted_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 词ID列表
	DeletedKeywordIds []int64 `json:"deleted_keyword_ids,omitempty" xml:"deleted_keyword_ids>int64,omitempty"`
}
