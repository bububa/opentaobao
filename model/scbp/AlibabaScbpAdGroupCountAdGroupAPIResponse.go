package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupcountadgroupAPIResponse 统计adgroup数量 API返回值
// alibaba.scbp.ad.group.count.ad.group
//
// 统计adgroup数量
type AlibabascbpadgroupcountadgroupAPIResponse struct {
	model.CommonResponse
	AlibabascbpadgroupcountadgroupAPIResponseModel
}

// AlibabascbpadgroupcountadgroupAPIResponseModel is 统计adgroup数量 成功返回结果
type AlibabascbpadgroupcountadgroupAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_count_ad_group_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
