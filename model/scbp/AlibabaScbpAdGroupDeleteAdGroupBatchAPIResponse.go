package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupDeleteAdGroupBatchAPIResponse 删除推广单元 API返回值
// alibaba.scbp.ad.group.delete.ad.group.batch
//
// 删除推广单元
type AlibabaScbpAdGroupDeleteAdGroupBatchAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdGroupDeleteAdGroupBatchAPIResponseModel
}

// AlibabaScbpAdGroupDeleteAdGroupBatchAPIResponseModel is 删除推广单元 成功返回结果
type AlibabaScbpAdGroupDeleteAdGroupBatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_delete_ad_group_batch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
