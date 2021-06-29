package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向销退入库单到仓结果回告 API请求
alibaba.ascp.uop.supplier.reverseorder.instorage.result

ERP回告销退入库单到仓信息回告
*/
type AlibabaAscpUopSupplierReverseorderInstorageResultRequest struct {
    model.Params
    // 消退入库单结果请求
    _instorageResultRequest   *Instorageresultrequest
}

// 初始化AlibabaAscpUopSupplierReverseorderInstorageResultRequest对象
func NewAlibabaAscpUopSupplierReverseorderInstorageResultRequest() *AlibabaAscpUopSupplierReverseorderInstorageResultRequest{
    return &AlibabaAscpUopSupplierReverseorderInstorageResultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierReverseorderInstorageResultRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.reverseorder.instorage.result"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierReverseorderInstorageResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InstorageResultRequest Setter
// 消退入库单结果请求
func (r *AlibabaAscpUopSupplierReverseorderInstorageResultRequest) SetInstorageResultRequest(_instorageResultRequest *Instorageresultrequest) error {
    r._instorageResultRequest = _instorageResultRequest
    r.Set("instorage_result_request", _instorageResultRequest)
    return nil
}

// InstorageResultRequest Getter
func (r AlibabaAscpUopSupplierReverseorderInstorageResultRequest) GetInstorageResultRequest() *Instorageresultrequest {
    return r._instorageResultRequest
}
