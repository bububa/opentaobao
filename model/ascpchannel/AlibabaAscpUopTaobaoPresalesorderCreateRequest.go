package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预售商家仓接单 APIRequest
alibaba.ascp.uop.taobao.presalesorder.create

预售商家仓接单
*/
type AlibabaAscpUopTaobaoPresalesorderCreateRequest struct {
    model.Params

    // 预售商家仓接单对象
    presalesOrderCreateRequest   *PresalesordercreaterequestTest 

}

func NewAlibabaAscpUopTaobaoPresalesorderCreateRequest() *AlibabaAscpUopTaobaoPresalesorderCreateRequest{
    return &AlibabaAscpUopTaobaoPresalesorderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopTaobaoPresalesorderCreateRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.taobao.presalesorder.create"
}

func (r AlibabaAscpUopTaobaoPresalesorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopTaobaoPresalesorderCreateRequest) SetPresalesOrderCreateRequest(presalesOrderCreateRequest *PresalesordercreaterequestTest) error {
    r.presalesOrderCreateRequest = presalesOrderCreateRequest
    r.Set("presales_order_create_request", presalesOrderCreateRequest)
    return nil
}

func (r AlibabaAscpUopTaobaoPresalesorderCreateRequest) GetPresalesOrderCreateRequest() *PresalesordercreaterequestTest {
    return r.presalesOrderCreateRequest
}

