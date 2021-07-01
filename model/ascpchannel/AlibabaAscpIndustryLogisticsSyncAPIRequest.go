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
type AlibabaAscpIndustryLogisticsSyncAPIRequest struct {
    model.Params
    // 参数
    _param   *LogisticsSyncSellerRequest
}

// 初始化AlibabaAscpIndustryLogisticsSyncAPIRequest对象
func NewAlibabaAscpIndustryLogisticsSyncRequest() *AlibabaAscpIndustryLogisticsSyncAPIRequest{
    return &AlibabaAscpIndustryLogisticsSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryLogisticsSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.industry.logistics.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryLogisticsSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 参数
func (r *AlibabaAscpIndustryLogisticsSyncAPIRequest) SetParam(_param *LogisticsSyncSellerRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaAscpIndustryLogisticsSyncAPIRequest) GetParam() *LogisticsSyncSellerRequest {
    return r._param
}
