package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除买赠活动 API请求
alibaba.wdk.marketing.itembuygift.deleteactivity

删除买赠活动
*/
type AlibabaWdkMarketingItembuygiftDeleteactivityRequest struct {
    model.Params
    // 要删除的活动信息
    _param   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItembuygiftDeleteactivityRequest对象
func NewAlibabaWdkMarketingItembuygiftDeleteactivityRequest() *AlibabaWdkMarketingItembuygiftDeleteactivityRequest{
    return &AlibabaWdkMarketingItembuygiftDeleteactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftDeleteactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itembuygift.deleteactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftDeleteactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 要删除的活动信息
func (r *AlibabaWdkMarketingItembuygiftDeleteactivityRequest) SetParam(_param *CommonActivityParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingItembuygiftDeleteactivityRequest) GetParam() *CommonActivityParam {
    return r._param
}
