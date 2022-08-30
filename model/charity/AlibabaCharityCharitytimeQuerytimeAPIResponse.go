package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeQuerytimeAPIResponse 外部查询公益时 API返回值
// alibaba.charity.charitytime.querytime
//
// 外部查询公益时
type AlibabaCharityCharitytimeQuerytimeAPIResponse struct {
	model.CommonResponse
	AlibabaCharityCharitytimeQuerytimeAPIResponseModel
}

// AlibabaCharityCharitytimeQuerytimeAPIResponseModel is 外部查询公益时 成功返回结果
type AlibabaCharityCharitytimeQuerytimeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_querytime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`
}
