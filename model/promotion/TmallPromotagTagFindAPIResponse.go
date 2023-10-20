package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallpromotagtagfindAPIResponse 查询标签接口 API返回值
// tmall.promotag.tag.find
//
// 查询用户创建的所有标签
type TmallpromotagtagfindAPIResponse struct {
	model.CommonResponse
	TmallpromotagtagfindAPIResponseModel
}

// TmallpromotagtagfindAPIResponseModel is 查询标签接口 成功返回结果
type TmallpromotagtagfindAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_promotag_tag_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果类型
	QueryResult *PromotionTagQuery `json:"query_result,omitempty" xml:"query_result,omitempty"`
}
