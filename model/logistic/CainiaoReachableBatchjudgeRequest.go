package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
是否派送可达判定批量查询接口 APIRequest
cainiao.reachable.batchjudge

提供给商家在发货之前做截单处理，输入物流商编码和收发货地址进行可达判定，目前支持国内主流的物流服务商, 支持快运和快递两种类型
*/
type CainiaoReachableBatchjudgeRequest struct {
    model.Params

    // 调用方对象
    clientInfo   *ClientInfoDto 

    // 1:快递 2:快运
    addressType   int64 

    // 收发信息
    data   *RoutingReachableBatchRequestDto 

}

func NewCainiaoReachableBatchjudgeRequest() *CainiaoReachableBatchjudgeRequest{
    return &CainiaoReachableBatchjudgeRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoReachableBatchjudgeRequest) GetApiMethodName() string {
    return "cainiao.reachable.batchjudge"
}

func (r CainiaoReachableBatchjudgeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoReachableBatchjudgeRequest) SetClientInfo(clientInfo *ClientInfoDto) error {
    r.clientInfo = clientInfo
    r.Set("client_info", clientInfo)
    return nil
}

func (r CainiaoReachableBatchjudgeRequest) GetClientInfo() *ClientInfoDto {
    return r.clientInfo
}

func (r *CainiaoReachableBatchjudgeRequest) SetAddressType(addressType int64) error {
    r.addressType = addressType
    r.Set("address_type", addressType)
    return nil
}

func (r CainiaoReachableBatchjudgeRequest) GetAddressType() int64 {
    return r.addressType
}

func (r *CainiaoReachableBatchjudgeRequest) SetData(data *RoutingReachableBatchRequestDto) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r CainiaoReachableBatchjudgeRequest) GetData() *RoutingReachableBatchRequestDto {
    return r.data
}

