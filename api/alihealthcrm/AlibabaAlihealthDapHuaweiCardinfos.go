package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
华为负一屏卡片查询 
alibaba.alihealth.dap.huawei.cardinfos

医疗健康频道卡片华为负一屏
*/
func AlibabaAlihealthDapHuaweiCardinfos(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthDapHuaweiCardinfosAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthDapHuaweiCardinfosAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthDapHuaweiCardinfosAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
