package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询税控设备加盘订购单详情 API请求
alibaba.einvoice.device.order.query

查询税控设备订购单详情
*/
type AlibabaEinvoiceDeviceOrderQueryRequest struct {
    model.Params
    // 税控设备订购单ID
    flowId   string
}

// 初始化AlibabaEinvoiceDeviceOrderQueryRequest对象
func NewAlibabaEinvoiceDeviceOrderQueryRequest() *AlibabaEinvoiceDeviceOrderQueryRequest{
    return &AlibabaEinvoiceDeviceOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceDeviceOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.einvoice.device.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceDeviceOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FlowId Setter
// 税控设备订购单ID
func (r *AlibabaEinvoiceDeviceOrderQueryRequest) SetFlowId(flowId string) error {
    r.flowId = flowId
    r.Set("flow_id", flowId)
    return nil
}

// FlowId Getter
func (r AlibabaEinvoiceDeviceOrderQueryRequest) GetFlowId() string {
    return r.flowId
}
