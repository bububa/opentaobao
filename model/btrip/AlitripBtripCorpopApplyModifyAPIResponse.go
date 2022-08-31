package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplyModifyAPIResponse 【商旅】修改出差审批单（行程） API返回值
// alitrip.btrip.corpop.apply.modify
//
// 【商旅】修改出差审批单（行程）
type AlitripBtripCorpopApplyModifyAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopApplyModifyAPIResponseModel
}

// AlitripBtripCorpopApplyModifyAPIResponseModel is 【商旅】修改出差审批单（行程） 成功返回结果
type AlitripBtripCorpopApplyModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
