package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomedynamicdatasubmitAPIResponse 提交小区动态信息 API返回值
// alibaba.alihouse.newhome.dynamicdata.submit
//
// 提交小区动态信息
type AlibabaalihousenewhomedynamicdatasubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomedynamicdatasubmitAPIResponseModel
}

// AlibabaalihousenewhomedynamicdatasubmitAPIResponseModel is 提交小区动态信息 成功返回结果
type AlibabaalihousenewhomedynamicdatasubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_dynamicdata_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomedynamicdatasubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
