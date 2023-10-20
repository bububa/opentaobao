package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordsdeleteAPIResponse 删除一批关键词 API返回值
// taobao.simba.keywords.delete
//
// 删除一批关键词
type TaobaosimbakeywordsdeleteAPIResponse struct {
	model.CommonResponse
	TaobaosimbakeywordsdeleteAPIResponseModel
}

// TaobaosimbakeywordsdeleteAPIResponseModel is 删除一批关键词 成功返回结果
type TaobaosimbakeywordsdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywords_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功删除的关键词列表
	Keywords []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
}
