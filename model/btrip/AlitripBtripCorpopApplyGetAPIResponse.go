package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplyGetAPIResponse 【商旅】查询审批单 API返回值
// alitrip.btrip.corpop.apply.get
//
// 【商旅】查询审批单
type AlitripBtripCorpopApplyGetAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopApplyGetAPIResponseModel
}

// AlitripBtripCorpopApplyGetAPIResponseModel is 【商旅】查询审批单 成功返回结果
type AlitripBtripCorpopApplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
