package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查看功能权限 APIResponse
taobao.simba.customers.sid.get

查询用户是否拥有某个功能权限
*/
type TaobaoSimbaCustomersSidGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCustomersSidGetResponse `json:"taobao_simba_customers_sid_get_response,omitempty"`
}

type TaobaoSimbaCustomersSidGetResponse struct {

    // 权限列表及是否有权限
    Result   *SidVo `json:"result,omitempty"`

}
