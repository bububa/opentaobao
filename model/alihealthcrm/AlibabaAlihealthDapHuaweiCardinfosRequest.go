package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
华为负一屏卡片查询 API请求
alibaba.alihealth.dap.huawei.cardinfos

医疗健康频道卡片华为负一屏
*/
type AlibabaAlihealthDapHuaweiCardinfosRequest struct {
    model.Params
    // source     HUAWEI_HAG,OPPO_OAG
    param   string
}

// 初始化AlibabaAlihealthDapHuaweiCardinfosRequest对象
func NewAlibabaAlihealthDapHuaweiCardinfosRequest() *AlibabaAlihealthDapHuaweiCardinfosRequest{
    return &AlibabaAlihealthDapHuaweiCardinfosRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDapHuaweiCardinfosRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dap.huawei.cardinfos"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDapHuaweiCardinfosRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// source     HUAWEI_HAG,OPPO_OAG
func (r *AlibabaAlihealthDapHuaweiCardinfosRequest) SetParam(param string) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaAlihealthDapHuaweiCardinfosRequest) GetParam() string {
    return r.param
}
