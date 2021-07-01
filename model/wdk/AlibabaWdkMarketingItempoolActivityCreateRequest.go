package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建活动新接口 API请求
alibaba.wdk.marketing.itempool.activity.create

创建活动新接口，支持新工具玩法
*/
type AlibabaWdkMarketingItempoolActivityCreateAPIRequest struct {
    model.Params
    // 创建活动请求入参
    _param   *ItemPoolActivity
}

// 初始化AlibabaWdkMarketingItempoolActivityCreateAPIRequest对象
func NewAlibabaWdkMarketingItempoolActivityCreateRequest() *AlibabaWdkMarketingItempoolActivityCreateAPIRequest{
    return &AlibabaWdkMarketingItempoolActivityCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolActivityCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.activity.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolActivityCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingItempoolActivityCreateAPIRequest) SetParam(_param *ItemPoolActivity) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingItempoolActivityCreateAPIRequest) GetParam() *ItemPoolActivity {
    return r._param
}
