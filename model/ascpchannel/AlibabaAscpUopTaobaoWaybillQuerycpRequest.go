package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询电子面单开放的CP列表 APIRequest
alibaba.ascp.uop.taobao.waybill.querycp

查询电子面单开放的CP列表
*/
type AlibabaAscpUopTaobaoWaybillQuerycpRequest struct {
    model.Params

    // 系统自动生成
    queryCpRequest   *Querycprequest 

}

func NewAlibabaAscpUopTaobaoWaybillQuerycpRequest() *AlibabaAscpUopTaobaoWaybillQuerycpRequest{
    return &AlibabaAscpUopTaobaoWaybillQuerycpRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopTaobaoWaybillQuerycpRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.taobao.waybill.querycp"
}

func (r AlibabaAscpUopTaobaoWaybillQuerycpRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopTaobaoWaybillQuerycpRequest) SetQueryCpRequest(queryCpRequest *Querycprequest) error {
    r.queryCpRequest = queryCpRequest
    r.Set("query_cp_request", queryCpRequest)
    return nil
}

func (r AlibabaAscpUopTaobaoWaybillQuerycpRequest) GetQueryCpRequest() *Querycprequest {
    return r.queryCpRequest
}

