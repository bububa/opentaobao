package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
门店改签合同接口 
alibaba.ele.fengniao.chainstore.contract.change

通过调用接口，门店改签物流服务包
*/
func AlibabaEleFengniaoChainstoreContractChange(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoChainstoreContractChangeRequest, session string) (*logistic.AlibabaEleFengniaoChainstoreContractChangeResponse, error) {
    var resp logistic.AlibabaEleFengniaoChainstoreContractChangeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
