package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询税控设备加盘订购单详情 API返回值 
alibaba.einvoice.device.order.query

查询税控设备订购单详情
*/
type AlibabaEinvoiceDeviceOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceDeviceOrderQueryResponse
}

// 查询税控设备加盘订购单详情 成功返回结果
type AlibabaEinvoiceDeviceOrderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_device_order_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
