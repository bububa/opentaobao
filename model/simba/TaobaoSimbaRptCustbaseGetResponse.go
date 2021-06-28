package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客户账户报表基础数据对象 APIResponse
taobao.simba.rpt.custbase.get

客户账户报表基础数据对象
*/
type TaobaoSimbaRptCustbaseGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_rpt_custbase_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 用户帐户结果
    
    RptCustBaseList   string `json:"rpt_cust_base_list,omitempty" xml:"