package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向销退入库单到仓结果回告 APIRequest
alibaba.ascp.uop.supplier.reverseorder.instorage.result

ERP回告销退入库单到仓信息回告
*/
type AlibabaAscpUopSupplierReverseorderInstorageResultRequest struct {
    model.Params

    // 消退入库单结果请求
    instorageResultRequest   *Instorageresultrequest 

}

func NewAlibabaAscpUopSupplierReverseorderInstorageResultRequest() *AlibabaAscpUopSupplierReverseorderInstorageResultRequest{
    return &AlibabaAscpUopSupplierReverseorderInstorageResultRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopSupplierReverseorderInstorageResultRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.reverseorder.instorage.result"
}

func (r AlibabaAscpUopSupplierReverseorderInstorageResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopSupplierReverseorderInstorageResultRequest) SetInstorageResultRequest(instorageResultRequest *Instorageresultrequest) error {
    r.instorageResultRequest = instorageResultRequest
    r.Set("instorage_result_request", instorageResultRequest)
    return nil
}

func (r AlibabaAscpUopSupplierReverseorderInstorageResultRequest) GetInstorageResultRequest() *Instorageresultrequest {
    return r.instorageResultRequest
}

