package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询今日消耗 API请求
alibaba.scbp.account.daycost.get

查询今日消耗
*/
type AlibabaScbpAccountDaycostGetRequest struct {
    model.Params
}

// 初始化AlibabaScbpAccountDaycostGetRequest对象
func NewAlibabaScbpAccountDaycostGetRequest() *AlibabaScbpAccountDaycostGetRequest{
    return &AlibabaScbpAccountDaycostGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountDaycostGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.account.daycost.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountDaycostGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
