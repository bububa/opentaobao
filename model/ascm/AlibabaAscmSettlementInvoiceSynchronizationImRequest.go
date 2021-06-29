package ascm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
英迈发票同步到结算 API请求
alibaba.ascm.settlement.invoice.synchronization.im

外部供应商通过此API将发货的发票信息推送给供应链中台结算系统
*/
type AlibabaAscmSettlementInvoiceSynchronizationImRequest struct {
    model.Params
    // im invoice xml
    xmlDataSlot   string
}

// 初始化AlibabaAscmSettlementInvoiceSynchronizationImRequest对象
func NewAlibabaAscmSettlementInvoiceSynchronizationImRequest() *AlibabaAscmSettlementInvoiceSynchronizationImRequest{
    return &AlibabaAscmSettlementInvoiceSynchronizationImRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscmSettlementInvoiceSynchronizationImRequest) GetApiMethodName() string {
    return "alibaba.ascm.settlement.invoice.synchronization.im"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscmSettlementInvoiceSynchronizationImRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// XmlDataSlot Setter
// im invoice xml
func (r *AlibabaAscmSettlementInvoiceSynchronizationImRequest) SetXmlDataSlot(xmlDataSlot string) error {
    r.xmlDataSlot = xmlDataSlot
    r.Set("xml_data_slot", xmlDataSlot)
    return nil
}

// XmlDataSlot Getter
func (r AlibabaAscmSettlementInvoiceSynchronizationImRequest) GetXmlDataSlot() string {
    return r.xmlDataSlot
}
