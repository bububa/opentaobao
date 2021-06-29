package ascm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
英迈发票同步到结算 APIRequest
alibaba.ascm.settlement.invoice.synchronization.im

外部供应商通过此API将发货的发票信息推送给供应链中台结算系统
*/
type AlibabaAscmSettlementInvoiceSynchronizationImRequest struct {
    model.Params

    // im invoice xml
    xmlDataSlot   string 

}

func NewAlibabaAscmSettlementInvoiceSynchronizationImRequest() *AlibabaAscmSettlementInvoiceSynchronizationImRequest{
    return &AlibabaAscmSettlementInvoiceSynchronizationImRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscmSettlementInvoiceSynchronizationImRequest) GetApiMethodName() string {
    return "alibaba.ascm.settlement.invoice.synchronization.im"
}

func (r AlibabaAscmSettlementInvoiceSynchronizationImRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscmSettlementInvoiceSynchronizationImRequest) SetXmlDataSlot(xmlDataSlot string) error {
    r.xmlDataSlot = xmlDataSlot
    r.Set("xml_data_slot", xmlDataSlot)
    return nil
}

func (r AlibabaAscmSettlementInvoiceSynchronizationImRequest) GetXmlDataSlot() string {
    return r.xmlDataSlot
}

