package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加商品池活动 API请求
alibaba.wdk.marketing.itempool.createactivity

添加商品池活动
*/
type AlibabaWdkMarketingItempoolCreateactivityRequest struct {
    model.Params
    // 创建活动请求入参
    _param   *ItemPoolActivity
}

// 初始化AlibabaWdkMarketingItempoolCreateactivityRequest对象
func NewAlibabaWdkMarketingItempoolCreateactivityRequest() *AlibabaWdkMarketingItempoolCreateactivityRequest{
    return &AlibabaWdkMarketingItempoolCreateactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolCreateactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.createactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolCreateactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingItempoolCreateactivityRequest) SetParam(_param *ItemPoolActivity) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingItempoolCreateactivityRequest) GetParam() *ItemPoolActivity {
    return r._param
}
