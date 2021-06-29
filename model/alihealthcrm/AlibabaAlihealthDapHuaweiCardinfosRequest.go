package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
华为负一屏卡片查询 APIRequest
alibaba.alihealth.dap.huawei.cardinfos

医疗健康频道卡片华为负一屏
*/
type AlibabaAlihealthDapHuaweiCardinfosRequest struct {
    model.Params

    // source     HUAWEI_HAG,OPPO_OAG
    param   string 

}

func NewAlibabaAlihealthDapHuaweiCardinfosRequest() *AlibabaAlihealthDapHuaweiCardinfosRequest{
    return &AlibabaAlihealthDapHuaweiCardinfosRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDapHuaweiCardinfosRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dap.huawei.cardinfos"
}

func (r AlibabaAlihealthDapHuaweiCardinfosRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDapHuaweiCardinfosRequest) SetParam(param string) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaAlihealthDapHuaweiCardinfosRequest) GetParam() string {
    return r.param
}

