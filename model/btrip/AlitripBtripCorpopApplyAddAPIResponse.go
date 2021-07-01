package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCorpopApplyAddAPIResponse
【商旅】isv添加审批单 API返回值
alitrip.btrip.corpop.apply.add

【商旅】isv添加审批单 */
type AlitripBtripCorpopApplyAddAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopApplyAddAPIResponseModel
}

// AlitripBtripCorpopApplyAddAPIResponseModel is 【商旅】isv添加审批单 成功返回结果
type AlitripBtripCorpopApplyAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参数
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
