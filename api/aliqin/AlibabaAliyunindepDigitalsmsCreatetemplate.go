package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
数字短信模板创建 
alibaba.aliyunindep.digitalsms.createtemplate

数字短信模板创建，给阿里云一方产品使用，类型：9
*/
func AlibabaAliyunindepDigitalsmsCreatetemplate(clt *core.SDKClient, req *aliqin.AlibabaAliyunindepDigitalsmsCreatetemplateRequest, session string) (*aliqin.AlibabaAliyunindepDigitalsmsCreatetemplateResponse, error) {
    var resp aliqin.AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
