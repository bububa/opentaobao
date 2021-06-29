package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流状态同步 APIRequest
alibaba.ascp.industry.logistics.sync

履约物流状态同步
*/
type AlibabaAscpIndustryLogisticsSyncRequest struct {
    model.Params

    // 参数
    param   *LogisticsSyncSellerRequest 

}

func NewAlibabaAscpIndustryLogisticsSyncRequest() *AlibabaAscpIndustryLogisticsSyncRequest{
    return &AlibabaAscpIndustryLogisticsSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpIndustryLogisticsSyncRequest) GetApiMethodName() string {
    return "alibaba.ascp.industry.logistics.sync"
}

func (r AlibabaAscpIndustryLogisticsSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpIndustryLogisticsSyncRequest) SetParam(param *LogisticsSyncSellerRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaAscpIndustryLogisticsSyncRequest) GetParam() *LogisticsSyncSellerRequest {
    return r.param
}

