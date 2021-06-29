package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询税控设备加盘订购单详情 APIRequest
alibaba.einvoice.device.order.query

查询税控设备订购单详情
*/
type AlibabaEinvoiceDeviceOrderQueryRequest struct {
    model.Params

    // 税控设备订购单ID
    flowId   string 

}

func NewAlibabaEinvoiceDeviceOrderQueryRequest() *AlibabaEinvoiceDeviceOrderQueryRequest{
    return &AlibabaEinvoiceDeviceOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceDeviceOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.einvoice.device.order.query"
}

func (r AlibabaEinvoiceDeviceOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceDeviceOrderQueryRequest) SetFlowId(flowId string) error {
    r.flowId = flowId
    r.Set("flow_id", flowId)
    return nil
}

func (r AlibabaEinvoiceDeviceOrderQueryRequest) GetFlowId() string {
    return r.flowId
}

