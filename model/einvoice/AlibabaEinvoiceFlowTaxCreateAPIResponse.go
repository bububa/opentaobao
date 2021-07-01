package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建税控开通工单 API返回值 
alibaba.einvoice.flow.tax.create

商户在业务前台订购税控产品后，调用阿里发票此接口，提交税号的入驻开通工单。此接口返回为工单的提交结果，非真正入驻结果。开通结果会在商户完成设备的部署安装 入驻完成后，由阿里发票通过消息异步通知到业务前台。
*/
type AlibabaEinvoiceFlowTaxCreateAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceFlowTaxCreateAPIResponseModel
}

// 创建税控开通工单 成功返回结果
type AlibabaEinvoiceFlowTaxCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_einvoice_flow_tax_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 工单ID，发票中台生成
    FlowId   string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
}
