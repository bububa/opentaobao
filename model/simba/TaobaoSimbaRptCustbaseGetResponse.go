package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
客户账户报表基础数据对象 APIResponse
taobao.simba.rpt.custbase.get

客户账户报表基础数据对象
*/
type TaobaoSimbaRptCustbaseGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRptCustbaseGetResponse `json:"taobao_simba_rpt_custbase_get_response,omitempty"`
}

type TaobaoSimbaRptCustbaseGetResponse struct {

    // 用户帐户结果
    RptCustBaseList   string `json:"rpt_cust_base_list,omitempty"`

}
