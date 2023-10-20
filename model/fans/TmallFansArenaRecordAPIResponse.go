package fans

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallfansarenarecordAPIResponse 记录完成擂台的用户 API返回值
// tmall.fans.arena.record
//
// 记录完成擂台的用户和完成分数
type TmallfansarenarecordAPIResponse struct {
	model.CommonResponse
	TmallfansarenarecordAPIResponseModel
}

// TmallfansarenarecordAPIResponseModel is 记录完成擂台的用户 成功返回结果
type TmallfansarenarecordAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fans_arena_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	FansResult *FansResult `json:"fans_result,omitempty" xml:"fans_result,omitempty"`
}
