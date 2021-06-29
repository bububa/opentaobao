package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易订单详情查询 API请求
alibaba.wdk.order.get

五道口三江单据查询接口
*/
type AlibabaWdkOrderGetRequest struct {
    model.Params
    // 系统自动生成
    idListQueryReq   *IdListQueryRequest
}

// 初始化AlibabaWdkOrderGetRequest对象
func NewAlibabaWdkOrderGetRequest() *AlibabaWdkOrderGetRequest{
    return &AlibabaWdkOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdListQueryReq Setter
// 系统自动生成
func (r *AlibabaWdkOrderGetRequest) SetIdListQueryReq(idListQueryReq *IdListQueryRequest) error {
    r.idListQueryReq = idListQueryReq
    r.Set("id_list_query_req", idListQueryReq)
    return nil
}

// IdListQueryReq Getter
func (r AlibabaWdkOrderGetRequest) GetIdListQueryReq() *IdListQueryRequest {
    return r.idListQueryReq
}
