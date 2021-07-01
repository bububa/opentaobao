package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
无线保镖SDK风控数据上报 
alibaba.security.jaq.wsgriskdata.report

无线保镖sdk根据用户的需要，上报数据到聚安全云端
*/
func AlibabaSecurityJaqWsgriskdataReport(clt *core.SDKClient, req *security.AlibabaSecurityJaqWsgriskdataReportAPIRequest, session string) (*security.AlibabaSecurityJaqWsgriskdataReportAPIResponse, error) {
    var resp security.AlibabaSecurityJaqWsgriskdataReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
