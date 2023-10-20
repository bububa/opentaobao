package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordsvonaddAPIResponse 创建一批关键词 API返回值
// taobao.simba.keywordsvon.add
//
// 创建一批关键词
type TaobaosimbakeywordsvonaddAPIResponse struct {
	model.CommonResponse
	TaobaosimbakeywordsvonaddAPIResponseModel
}

// TaobaosimbakeywordsvonaddAPIResponseModel is 创建一批关键词 成功返回结果
type TaobaosimbakeywordsvonaddAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywordsvon_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词列表
	Keywords []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
}
