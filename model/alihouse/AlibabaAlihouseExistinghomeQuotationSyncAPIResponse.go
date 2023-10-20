package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomequotationsyncAPIResponse 二手房行情数据同步 API返回值
// alibaba.alihouse.existinghome.quotation.sync
//
// 二手房行情数据同步
type AlibabaalihouseexistinghomequotationsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomequotationsyncAPIResponseModel
}

// AlibabaalihouseexistinghomequotationsyncAPIResponseModel is 二手房行情数据同步 成功返回结果
type AlibabaalihouseexistinghomequotationsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_quotation_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaalihouseexistinghomequotationsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
