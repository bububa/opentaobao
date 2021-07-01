package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
薄荷同步三方记录 
alibaba.fmhealth.weight.lossplan.syncweightdata

用于三方薄荷同步数据到健康会员
*/
func AlibabaFmhealthWeightLossplanSyncweightdata(clt *core.SDKClient, req *alihealth2.AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest, session string) (*alihealth2.AlibabaFmhealthWeightLossplanSyncweightdataAPIResponse, error) {
    var resp alihealth2.AlibabaFmhealthWeightLossplanSyncweightdataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
