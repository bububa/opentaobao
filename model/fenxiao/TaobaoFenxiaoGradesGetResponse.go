package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分销商等级查询 APIResponse
taobao.fenxiao.grades.get

根据供应商ID，查询他的分销商等级信息
*/
type TaobaoFenxiaoGradesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoGradesGetResponse `json:"fenxiao_grades_get_response,omitempty"` 
    TaobaoFenxiaoGradesGetResponse
}

/* model for simplify = false
type TaobaoFenxiaoGradesGetResponse struct {

    // 分销商等级信息
    
    FenxiaoGrades  struct {
        FenxiaoGrade  []FenxiaoGrade `json:"fenxiao_grade,omitempty"`
    } `json:"fenxiao_grades,omitempty"`
    

}
*/

type TaobaoFenxiaoGradesGetResponse struct {

    // 分销商等级信息
    FenxiaoGrades   []FenxiaoGrade `json:"fenxiao_grades,omitempty"`

}
