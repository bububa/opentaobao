package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生成发布使用的版本号 API请求
alibaba.wdk.marketing.version.generate

生成发布使用的版本号
*/
type AlibabaWdkMarketingVersionGenerateRequest struct {
    model.Params
    // 档期版本号参数信息
    _param   *SeasonVersionParam
}

// 初始化AlibabaWdkMarketingVersionGenerateRequest对象
func NewAlibabaWdkMarketingVersionGenerateRequest() *AlibabaWdkMarketingVersionGenerateRequest{
    return &AlibabaWdkMarketingVersionGenerateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingVersionGenerateRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.version.generate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingVersionGenerateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 档期版本号参数信息
func (r *AlibabaWdkMarketingVersionGenerateRequest) SetParam(_param *SeasonVersionParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingVersionGenerateRequest) GetParam() *SeasonVersionParam {
    return r._param
}
