package nazca

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
法庭提交和结案案件量接口 API返回值 
alibaba.infodept.lassen.casestatistics.get

功能描述：云嘉为浙江省高院制作数据大屏，需展示网上法庭相关数据，我方为省高院提供浙江省内法院收案和结案的案件量，开放数据接口，供其调取这两组数据。
*/
type AlibabaInfodeptLassenCasestatisticsGetAPIResponse struct {
    model.CommonResponse
    AlibabaInfodeptLassenCasestatisticsGetAPIResponseModel
}

// 法庭提交和结案案件量接口 成功返回结果
type AlibabaInfodeptLassenCasestatisticsGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_infodept_lassen_casestatistics_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
