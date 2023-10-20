package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosFalconPosCounterQueryAPIResponse 云POS查看专柜属性 API返回值
// alibaba.mos.falcon.pos.counter.query
//
// 银泰商业获取专柜是否支持小数等属性查看
type AlibabaMosFalconPosCounterQueryAPIResponse struct {
	model.CommonResponse
	AlibabaMosFalconPosCounterQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosFalconPosCounterQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosFalconPosCounterQueryAPIResponseModel).Reset()
}

// AlibabaMosFalconPosCounterQueryAPIResponseModel is 云POS查看专柜属性 成功返回结果
type AlibabaMosFalconPosCounterQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_falcon_pos_counter_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMosFalconPosCounterQueryResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosFalconPosCounterQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosFalconPosCounterQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosFalconPosCounterQueryAPIResponse)
	},
}

// GetAlibabaMosFalconPosCounterQueryAPIResponse 从 sync.Pool 获取 AlibabaMosFalconPosCounterQueryAPIResponse
func GetAlibabaMosFalconPosCounterQueryAPIResponse() *AlibabaMosFalconPosCounterQueryAPIResponse {
	return poolAlibabaMosFalconPosCounterQueryAPIResponse.Get().(*AlibabaMosFalconPosCounterQueryAPIResponse)
}

// ReleaseAlibabaMosFalconPosCounterQueryAPIResponse 将 AlibabaMosFalconPosCounterQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosFalconPosCounterQueryAPIResponse(v *AlibabaMosFalconPosCounterQueryAPIResponse) {
	v.Reset()
	poolAlibabaMosFalconPosCounterQueryAPIResponse.Put(v)
}
