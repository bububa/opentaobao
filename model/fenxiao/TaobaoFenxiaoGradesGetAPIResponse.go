package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoGradesGetAPIResponse 分销商等级查询 API返回值
// taobao.fenxiao.grades.get
//
// 根据供应商ID，查询他的分销商等级信息
type TaobaoFenxiaoGradesGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoGradesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoGradesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoGradesGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoGradesGetAPIResponseModel is 分销商等级查询 成功返回结果
type TaobaoFenxiaoGradesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_grades_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分销商等级信息
	FenxiaoGrades []FenxiaoGrade `json:"fenxiao_grades,omitempty" xml:"fenxiao_grades>fenxiao_grade,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoGradesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FenxiaoGrades = m.FenxiaoGrades[:0]
}

var poolTaobaoFenxiaoGradesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoGradesGetAPIResponse)
	},
}

// GetTaobaoFenxiaoGradesGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoGradesGetAPIResponse
func GetTaobaoFenxiaoGradesGetAPIResponse() *TaobaoFenxiaoGradesGetAPIResponse {
	return poolTaobaoFenxiaoGradesGetAPIResponse.Get().(*TaobaoFenxiaoGradesGetAPIResponse)
}

// ReleaseTaobaoFenxiaoGradesGetAPIResponse 将 TaobaoFenxiaoGradesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoGradesGetAPIResponse(v *TaobaoFenxiaoGradesGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoGradesGetAPIResponse.Put(v)
}
