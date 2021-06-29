package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家修改运单号 APIRequest
alibaba.ascp.uop.supplier.consignorder.notify.tms.change

供应商可以通过此接口，对出库回告上报的运单号进行修改，目前一次调用只能支持一个运单号的修改
*/
type AlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest struct {
    model.Params

    // 修改运单号请求模型
    modifyMailNoRequest   *Modifymailnorequest 

}

func NewAlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest() *AlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest{
    return &AlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.consignorder.notify.tms.change"
}

func (r AlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest) SetModifyMailNoRequest(modifyMailNoRequest *Modifymailnorequest) error {
    r.modifyMailNoRequest = modifyMailNoRequest
    r.Set("modify_mail_no_request", modifyMailNoRequest)
    return nil
}

func (r AlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest) GetModifyMailNoRequest() *Modifymailnorequest {
    return r.modifyMailNoRequest
}

