package nazca

import (
    "github.com/bububa/opentaobao/model"
)

/* 
法庭提交和结案案件量接口 APIResponse
alibaba.infodept.lassen.casestatistics.get

功能描述：云嘉为浙江省高院制作数据大屏，需展示网上法庭相关数据，我方为省高院提供浙江省内法院收案和结案的案件量，开放数据接口，供其调取这两组数据。
*/
type AlibabaInfodeptLassenCasestatisticsGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInfodeptLassenCasestatisticsGetResponse `json:"alibaba_infodept_lassen_casestatistics_get_response,omitempty"` 
    AlibabaInfodeptLassenCasestatisticsGetResponse
}

/* model for simplify = false
type AlibabaInfodeptLassenCasestatisticsGetResponse struct {

    // result
    
    Result   string `json:"result,omitempty"`
    

    // result
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInfodeptLassenCasestatisticsGetResponse struct {

    // result
    Result   string `json:"result,omitempty"`

    // result
    Result   string `json:"result,omitempty"`

}
