package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopapplysearchAPIResponse 【商旅】搜索审批单列表 API返回值
// alitrip.btrip.corpop.apply.search
//
// 【商旅】搜索审批单列表
type AlitripbtripcorpopapplysearchAPIResponse struct {
	model.CommonResponse
	AlitripbtripcorpopapplysearchAPIResponseModel
}

// AlitripbtripcorpopapplysearchAPIResponseModel is 【商旅】搜索审批单列表 成功返回结果
type AlitripbtripcorpopapplysearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
