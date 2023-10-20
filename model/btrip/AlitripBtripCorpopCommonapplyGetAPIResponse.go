package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopCommonapplyGetAPIResponse 商旅审批单通用查询接口 API返回值
// alitrip.btrip.corpop.commonapply.get
//
// 商旅审批单通用查询接口
type AlitripBtripCorpopCommonapplyGetAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopCommonapplyGetAPIResponseModel
}

// AlitripBtripCorpopCommonapplyGetAPIResponseModel is 商旅审批单通用查询接口 成功返回结果
type AlitripBtripCorpopCommonapplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_commonapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
