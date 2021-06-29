package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建全场活动 API请求
alibaba.wdk.marketing.fullrange.createactivity

创建全场活动
*/
type AlibabaWdkMarketingFullrangeCreateactivityRequest struct {
    model.Params
    // 创建活动请求入参
    _param   *FullRangeActivity
}

// 初始化AlibabaWdkMarketingFullrangeCreateactivityRequest对象
func NewAlibabaWdkMarketingFullrangeCreateactivityRequest() *AlibabaWdkMarketingFullrangeCreateactivityRequest{
    return &AlibabaWdkMarketingFullrangeCreateactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeCreateactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.createactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeCreateactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingFullrangeCreateactivityRequest) SetParam(_param *FullRangeActivity) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingFullrangeCreateactivityRequest) GetParam() *FullRangeActivity {
    return r._param
}
