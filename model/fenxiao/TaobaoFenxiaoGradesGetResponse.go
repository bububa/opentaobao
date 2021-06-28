package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商等级查询 APIResponse
taobao.fenxiao.grades.get

根据供应商ID，查询他的分销商等级信息
*/
type TaobaoFenxiaoGradesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_grades_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分销商等级信息
    
    FenxiaoGrades   []FenxiaoGrade `json:"fenxiao_grades,omitempty" xml:"