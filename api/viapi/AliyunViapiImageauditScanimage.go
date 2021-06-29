package viapi

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/viapi"
)

/* 
绿网-内容安全 
aliyun.viapi.imageaudit.scanimage

绿网-内容安全技术是基于阿里云视觉分析技术和深度识别技术，并经过在阿里经济体内和云上客户的多领域、多场景的广泛应用和不断优化，可提供风险和治理领域的图像识别、定位、检索等全面服务能力，不仅可以降低色情、涉恐、涉政、广告、垃圾信息等违规风险，而且能大幅度降低人工审核成本。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
func AliyunViapiImageauditScanimage(clt *core.SDKClient, req *viapi.AliyunViapiImageauditScanimageRequest, session string) (*viapi.AliyunViapiImageauditScanimageAPIResponse, error) {
    var resp viapi.AliyunViapiImageauditScanimageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
