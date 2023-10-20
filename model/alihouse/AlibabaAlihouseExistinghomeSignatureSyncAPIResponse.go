package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomesignaturesyncAPIResponse 二手房电子签章数据同步 API返回值
// alibaba.alihouse.existinghome.signature.sync
//
// 二手房电子签章数据同步
type AlibabaalihouseexistinghomesignaturesyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomesignaturesyncAPIResponseModel
}

// AlibabaalihouseexistinghomesignaturesyncAPIResponseModel is 二手房电子签章数据同步 成功返回结果
type AlibabaalihouseexistinghomesignaturesyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_signature_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomesignaturesyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
