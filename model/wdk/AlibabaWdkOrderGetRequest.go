package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易订单详情查询 APIRequest
alibaba.wdk.order.get

五道口三江单据查询接口
*/
type AlibabaWdkOrderGetRequest struct {
    model.Params

    // 系统自动生成
    idListQueryReq   *IdListQueryRequest 

}

func NewAlibabaWdkOrderGetRequest() *AlibabaWdkOrderGetRequest{
    return &AlibabaWdkOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkOrderGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.get"
}

func (r AlibabaWdkOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkOrderGetRequest) SetIdListQueryReq(idListQueryReq *IdListQueryRequest) error {
    r.idListQueryReq = idListQueryReq
    r.Set("id_list_query_req", idListQueryReq)
    return nil
}

func (r AlibabaWdkOrderGetRequest) GetIdListQueryReq() *IdListQueryRequest {
    return r.idListQueryReq
}

