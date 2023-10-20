package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychaintraincityAPIResponse 火车站数据查询 API返回值
// alitrip.btrip.supplychain.train.city
//
// 火车站数据查询
type AlitripbtripsupplychaintraincityAPIResponse struct {
	model.CommonResponse
	AlitripbtripsupplychaintraincityAPIResponseModel
}

// AlitripbtripsupplychaintraincityAPIResponseModel is 火车站数据查询 成功返回结果
type AlitripbtripsupplychaintraincityAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_train_city_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果码
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果对象
	Result *OpenApiSuggestRs `json:"result,omitempty" xml:"result,omitempty"`
	// 结果信息
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
