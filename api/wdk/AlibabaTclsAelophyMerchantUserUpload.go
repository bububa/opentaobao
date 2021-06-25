package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商家会员数据上传 
alibaba.tcls.aelophy.merchant.user.upload

商家会员数据上传
*/
func AlibabaTclsAelophyMerchantUserUpload(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantUserUploadRequest, session string) (*wdk.AlibabaTclsAelophyMerchantUserUploadResponse, error) {
    var resp wdk.AlibabaTclsAelophyMerchantUserUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
