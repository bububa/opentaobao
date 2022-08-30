package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectTradeitemAPIResponse 新房品同步 API返回值
// alibaba.alihouse.newhome.project.tradeitem
//
// 新房品同步
type AlibabaAlihouseNewhomeProjectTradeitemAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectTradeitemAPIResponseModel
}

// AlibabaAlihouseNewhomeProjectTradeitemAPIResponseModel is 新房品同步 成功返回结果
type AlibabaAlihouseNewhomeProjectTradeitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_tradeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectTradeitemResult `json:"result,omitempty" xml:"result,omitempty"`
}
