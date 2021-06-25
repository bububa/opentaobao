package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
物流信息回传接口 
alibaba.alihealth.nr.delivery.history.save

商家ERP回传物流信息
*/
func AlibabaAlihealthNrDeliveryHistorySave(clt *core.SDKClient, req *logistic.AlibabaAlihealthNrDeliveryHistorySaveRequest, session string) (*logistic.AlibabaAlihealthNrDeliveryHistorySaveResponse, error) {
    var resp logistic.AlibabaAlihealthNrDeliveryHistorySaveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
