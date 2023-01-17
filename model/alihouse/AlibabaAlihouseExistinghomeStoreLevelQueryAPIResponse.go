package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse 门店等级评分查询 API返回值
// alibaba.alihouse.existinghome.store.level.query
//
// 门店等级评分查询
type AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponseModel
}

// AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponseModel is 门店等级评分查询 成功返回结果
type AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_store_level_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AlibabaAlihouseExistinghomeStoreLevelQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
