package drug

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drug"
)

/* 
阿里健康医保支付信息获取 
alibaba.alihealth.nr.trade.medical.insurance.get

阿里健康医保支付信息获取
*/
func AlibabaAlihealthNrTradeMedicalInsuranceGet(clt *core.SDKClient, req *drug.AlibabaAlihealthNrTradeMedicalInsuranceGetRequest, session string) (*drug.AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse, error) {
    var resp drug.AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
