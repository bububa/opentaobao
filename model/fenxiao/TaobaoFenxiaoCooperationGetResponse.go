package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商或分销商获取合作关系信息 APIResponse
taobao.fenxiao.cooperation.get

获取供应商的合作关系信息
*/
type TaobaoFenxiaoCooperationGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoCooperationGetResponse `json:"fenxiao_cooperation_get_response,omitempty"` 
    TaobaoFenxiaoCooperationGetResponse
}

/* model for simplify = false
type TaobaoFenxiaoCooperationGetResponse struct {

    // 结果记录数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // 合作分销关系
    
    Cooperations  struct {
        Cooperation  []Cooperation `json:"cooperation,omitempty"`
    } `json:"cooperations,omitempty"`
    

}
*/

type TaobaoFenxiaoCooperationGetResponse struct {

    // 结果记录数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 合作分销关系
    Cooperations   []Cooperation `json:"cooperations,omitempty"`

}
