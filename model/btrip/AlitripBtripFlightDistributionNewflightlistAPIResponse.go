package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionnewflightlistAPIResponse 商旅机票航班列表接口，用于分销询价V2 API返回值
// alitrip.btrip.flight.distribution.newflightlist
//
// 商旅机票航班列表接口，用于分销询价V2
type AlitripbtripflightdistributionnewflightlistAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionnewflightlistAPIResponseModel
}

// AlitripbtripflightdistributionnewflightlistAPIResponseModel is 商旅机票航班列表接口，用于分销询价V2 成功返回结果
type AlitripbtripflightdistributionnewflightlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_newflightlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
