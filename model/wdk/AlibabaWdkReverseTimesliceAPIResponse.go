package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkreversetimesliceAPIResponse 逆向取货时间片查询 API返回值
// alibaba.wdk.reverse.timeslice
//
// 逆向取货时间片查询
type AlibabawdkreversetimesliceAPIResponse struct {
	model.CommonResponse
	AlibabawdkreversetimesliceAPIResponseModel
}

// AlibabawdkreversetimesliceAPIResponseModel is 逆向取货时间片查询 成功返回结果
type AlibabawdkreversetimesliceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_timeslice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}
