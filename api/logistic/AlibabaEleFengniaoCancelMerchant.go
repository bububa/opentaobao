package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
商户取消 
alibaba.ele.fengniao.cancel.merchant

商户取消配送
*/
func AlibabaEleFengniaoCancelMerchant(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoCancelMerchantAPIRequest, session string) (*logistic.AlibabaEleFengniaoCancelMerchantAPIResponse, error) {
    var resp logistic.AlibabaEleFengniaoCancelMerchantAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
