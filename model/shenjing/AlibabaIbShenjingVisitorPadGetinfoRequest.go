package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取OSS上传参数 APIRequest
alibaba.ib.shenjing.visitor.pad.getinfo

PAD 端获取OSS上传参数，向OSS服务器上传图片。
*/
type AlibabaIbShenjingVisitorPadGetinfoRequest struct {
    model.Params

}

func NewAlibabaIbShenjingVisitorPadGetinfoRequest() *AlibabaIbShenjingVisitorPadGetinfoRequest{
    return &AlibabaIbShenjingVisitorPadGetinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIbShenjingVisitorPadGetinfoRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.getinfo"
}

func (r AlibabaIbShenjingVisitorPadGetinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


