package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全获取异步图文识别结果接口 
alibaba.security.jaq.ocr.image.async.detect.results.fetch

获取异步图像字符识别结果接口根据图像检测接口返回taskid来获取对应图像的检测结果
*/
func AlibabaSecurityJaqOcrImageAsyncDetectResultsFetch(clt *core.SDKClient, req *security.AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest, session string) (*security.AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse, error) {
    var resp security.AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
