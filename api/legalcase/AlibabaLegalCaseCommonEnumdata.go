package legalcase

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalcase"
)

/* 
获取通用枚举值接口 
alibaba.legal.case.common.enumdata

获取通用枚举值接口
*/
func AlibabaLegalCaseCommonEnumdata(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseCommonEnumdataRequest, session string) (*legalcase.AlibabaLegalCaseCommonEnumdataAPIResponse, error) {
    var resp legalcase.AlibabaLegalCaseCommonEnumdataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
