package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeActivityQueryAPIResponse 五福活动经纪人获取 API返回值
// alibaba.alihouse.existinghome.activity.query
//
// 五福活动经纪人获取
type AlibabaAlihouseExistinghomeActivityQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeActivityQueryAPIResponseModel
}

// AlibabaAlihouseExistinghomeActivityQueryAPIResponseModel is 五福活动经纪人获取 成功返回结果
type AlibabaAlihouseExistinghomeActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihouseExistinghomeActivityQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
