package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomequeryhousetaginfoAPIResponse 活动标查询接口 API返回值
// alibaba.alihouse.existinghome.query.house.tag.info
//
// 活动标查询接口
type AlibabaalihouseexistinghomequeryhousetaginfoAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomequeryhousetaginfoAPIResponseModel
}

// AlibabaalihouseexistinghomequeryhousetaginfoAPIResponseModel is 活动标查询接口 成功返回结果
type AlibabaalihouseexistinghomequeryhousetaginfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_query_house_tag_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaalihouseexistinghomequeryhousetaginfoResult `json:"result,omitempty" xml:"result,omitempty"`
}
