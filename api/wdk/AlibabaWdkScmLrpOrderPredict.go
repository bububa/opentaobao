package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
单量预测 
alibaba.wdk.scm.lrp.order.predict

提供基于门店和配送站的订单量预测，可用于提前安排人力资源
*/
func AlibabaWdkScmLrpOrderPredict(clt *core.SDKClient, req *wdk.AlibabaWdkScmLrpOrderPredictRequest, session string) (*wdk.AlibabaWdkScmLrpOrderPredictResponse, error) {
    var resp wdk.AlibabaWdkScmLrpOrderPredictAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
