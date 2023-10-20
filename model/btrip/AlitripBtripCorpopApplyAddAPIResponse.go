package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopapplyaddAPIResponse 【商旅】isv添加审批单 API返回值
// alitrip.btrip.corpop.apply.add
//
// 【商旅】isv添加审批单
type AlitripbtripcorpopapplyaddAPIResponse struct {
	model.CommonResponse
	AlitripbtripcorpopapplyaddAPIResponseModel
}

// AlitripbtripcorpopapplyaddAPIResponseModel is 【商旅】isv添加审批单 成功返回结果
type AlitripbtripcorpopapplyaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参数
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
