package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款详情 API请求
alibaba.wdk.reverse.reversedetail

退款详情
*/
type AlibabaWdkReverseReversedetailRequest struct {
    model.Params
    // 退款单id
    reverseId   string
}

// 初始化AlibabaWdkReverseReversedetailRequest对象
func NewAlibabaWdkReverseReversedetailRequest() *AlibabaWdkReverseReversedetailRequest{
    return &AlibabaWdkReverseReversedetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseReversedetailRequest) GetApiMethodName() string {
    return "alibaba.wdk.reverse.reversedetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseReversedetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReverseId Setter
// 退款单id
func (r *AlibabaWdkReverseReversedetailRequest) SetReverseId(reverseId string) error {
    r.reverseId = reverseId
    r.Set("reverse_id", reverseId)
    return nil
}

// ReverseId Getter
func (r AlibabaWdkReverseReversedetailRequest) GetReverseId() string {
    return r.reverseId
}
