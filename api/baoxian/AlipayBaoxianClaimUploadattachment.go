package baoxian

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baoxian"
)

/* 
资料上传接口 
alipay.baoxian.claim.uploadattachment

给合作伙伴上传申请理赔材料
*/
func AlipayBaoxianClaimUploadattachment(clt *core.SDKClient, req *baoxian.AlipayBaoxianClaimUploadattachmentRequest, session string) (*baoxian.AlipayBaoxianClaimUploadattachmentResponse, error) {
    var resp baoxian.AlipayBaoxianClaimUploadattachmentAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
