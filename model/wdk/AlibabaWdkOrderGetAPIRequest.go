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
type AlibabaWdkOrderGetAPIRequest struct {
    model.Params
    // 系统自动生成
    _idListQueryReq   *IdListQueryRequest
}

// 初始化AlibabaWdkOrderGetAPIRequest对象
func NewAlibabaWdkOrderGetRequest() *AlibabaWdkOrderGetAPIRequest{
    return &AlibabaWdkOrderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdListQueryReq Setter
// 系统自动生成
func (r *AlibabaWdkOrderGetAPIRequest) SetIdListQueryReq(_idListQueryReq *IdListQueryRequest) error {
    r._idListQueryReq = _idListQueryReq
    r.Set("id_list_query_req", _idListQueryReq)
    return nil
}

// IdListQueryReq Getter
func (r AlibabaWdkOrderGetAPIRequest) GetIdListQueryReq() *IdListQueryRequest {
    return r._idListQueryReq
}
