package legalcase

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalcase"
)

/* 
消息通知 
alibaba.legal.case.common.notice

同步通知给诉讼系统
*/
func AlibabaLegalCaseCommonNotice(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseCommonNoticeAPIRequest, session string) (*legalcase.AlibabaLegalCaseCommonNoticeAPIResponse, error) {
    var resp legalcase.AlibabaLegalCaseCommonNoticeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
