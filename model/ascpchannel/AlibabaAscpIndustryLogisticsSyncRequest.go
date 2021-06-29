package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流状态同步 API请求
alibaba.ascp.industry.logistics.sync

履约物流状态同步
*/
type AlibabaAscpIndustryLogisticsSyncRequest struct {
    model.Params
    // 参数
    _param   *LogisticsSyncSellerRequest
}

// 初始化AlibabaAscpIndustryLogisticsSyncRequest对象
func NewAlibabaAscpIndustryLogisticsSyncRequest() *AlibabaAscpIndustryLogisticsSyncRequest{
    return &AlibabaAscpIndustryLogisticsSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryLogisticsSyncRequest) GetApiMethodName() string {
    return "alibaba.ascp.industry.logistics.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryLogisticsSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 参数
func (r *AlibabaAscpIndustryLogisticsSyncRequest) SetParam(_param *LogisticsSyncSellerRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaAscpIndustryLogisticsSyncRequest) GetParam() *LogisticsSyncSellerRequest {
    return r._param
}
