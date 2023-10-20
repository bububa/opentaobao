package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCrowdDmpCrowdBindAPIResponse 直通车绑定达摩盘人群 API返回值
// taobao.subway.crowd.dmp.crowd.bind
//
// 直通车绑定达摩盘人群
type TaobaoSubwayCrowdDmpCrowdBindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayCrowdDmpCrowdBindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayCrowdDmpCrowdBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayCrowdDmpCrowdBindAPIResponseModel).Reset()
}

// TaobaoSubwayCrowdDmpCrowdBindAPIResponseModel is 直通车绑定达摩盘人群 成功返回结果
type TaobaoSubwayCrowdDmpCrowdBindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_crowd_dmp_crowd_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 绑定结果数据
	Result []CrowdBindTopResultDto `json:"result,omitempty" xml:"result>crowd_bind_top_result_dto,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubwayCrowdDmpCrowdBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.Message = ""
	m.ResultCode = nil
}

var poolTaobaoSubwayCrowdDmpCrowdBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayCrowdDmpCrowdBindAPIResponse)
	},
}

// GetTaobaoSubwayCrowdDmpCrowdBindAPIResponse 从 sync.Pool 获取 TaobaoSubwayCrowdDmpCrowdBindAPIResponse
func GetTaobaoSubwayCrowdDmpCrowdBindAPIResponse() *TaobaoSubwayCrowdDmpCrowdBindAPIResponse {
	return poolTaobaoSubwayCrowdDmpCrowdBindAPIResponse.Get().(*TaobaoSubwayCrowdDmpCrowdBindAPIResponse)
}

// ReleaseTaobaoSubwayCrowdDmpCrowdBindAPIResponse 将 TaobaoSubwayCrowdDmpCrowdBindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayCrowdDmpCrowdBindAPIResponse(v *TaobaoSubwayCrowdDmpCrowdBindAPIResponse) {
	v.Reset()
	poolTaobaoSubwayCrowdDmpCrowdBindAPIResponse.Put(v)
}
