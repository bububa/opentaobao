package westcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmGradeGetAPIResponse 获取等级列表 API返回值
// alibaba.westcrm.grade.get
//
// 获取会员卡等级列表
type AlibabaWestcrmGradeGetAPIResponse struct {
	model.CommonResponse
	AlibabaWestcrmGradeGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWestcrmGradeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWestcrmGradeGetAPIResponseModel).Reset()
}

// AlibabaWestcrmGradeGetAPIResponseModel is 获取等级列表 成功返回结果
type AlibabaWestcrmGradeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_westcrm_grade_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWestcrmGradeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWestcrmGradeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWestcrmGradeGetAPIResponse)
	},
}

// GetAlibabaWestcrmGradeGetAPIResponse 从 sync.Pool 获取 AlibabaWestcrmGradeGetAPIResponse
func GetAlibabaWestcrmGradeGetAPIResponse() *AlibabaWestcrmGradeGetAPIResponse {
	return poolAlibabaWestcrmGradeGetAPIResponse.Get().(*AlibabaWestcrmGradeGetAPIResponse)
}

// ReleaseAlibabaWestcrmGradeGetAPIResponse 将 AlibabaWestcrmGradeGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWestcrmGradeGetAPIResponse(v *AlibabaWestcrmGradeGetAPIResponse) {
	v.Reset()
	poolAlibabaWestcrmGradeGetAPIResponse.Put(v)
}
