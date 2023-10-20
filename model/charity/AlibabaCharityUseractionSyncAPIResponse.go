package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharityuseractionsyncAPIResponse 用户公益行为同步 API返回值
// alibaba.charity.useraction.sync
//
// 外部公益活动，用户公益行为同步
type AlibabacharityuseractionsyncAPIResponse struct {
	model.CommonResponse
	AlibabacharityuseractionsyncAPIResponseModel
}

// AlibabacharityuseractionsyncAPIResponseModel is 用户公益行为同步 成功返回结果
type AlibabacharityuseractionsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_useraction_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`
}
