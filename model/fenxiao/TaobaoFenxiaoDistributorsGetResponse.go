package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取分销商信息 APIResponse
taobao.fenxiao.distributors.get

查询和当前登录供应商有合作关系的分销商的信息
*/
type TaobaoFenxiaoDistributorsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoDistributorsGetResponse `json:"fenxiao_distributors_get_response,omitempty"` 
    TaobaoFenxiaoDistributorsGetResponse
}

/* model for simplify = false
type TaobaoFenxiaoDistributorsGetResponse struct {

    // 分销商详细信息
    
    Distributors  struct {
        Distributor  []Distributor `json:"distributor,omitempty"`
    } `json:"distributors,omitempty"`
    

}
*/

type TaobaoFenxiaoDistributorsGetResponse struct {

    // 分销商详细信息
    Distributors   []Distributor `json:"distributors,omitempty"`

}
