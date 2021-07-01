package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoGradesGetAPIResponse
分销商等级查询 API返回值
taobao.fenxiao.grades.get

根据供应商ID，查询他的分销商等级信息 */
type TaobaoFenxiaoGradesGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoGradesGetAPIResponseModel
}

// TaobaoFenxiaoGradesGetAPIResponseModel is 分销商等级查询 成功返回结果
type TaobaoFenxiaoGradesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_grades_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分销商等级信息
	FenxiaoGrades []FenxiaoGrade `json:"fenxiao_grades,omitempty" xml:"fenxiao_grades>fenxiao_grade,omitempty"`
}
