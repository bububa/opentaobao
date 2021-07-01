package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取OSS上传参数 API请求
alibaba.ib.shenjing.visitor.pad.getinfo

PAD 端获取OSS上传参数，向OSS服务器上传图片。
*/
type AlibabaIbShenjingVisitorPadGetinfoAPIRequest struct {
    model.Params
}

// 初始化AlibabaIbShenjingVisitorPadGetinfoAPIRequest对象
func NewAlibabaIbShenjingVisitorPadGetinfoRequest() *AlibabaIbShenjingVisitorPadGetinfoAPIRequest{
    return &AlibabaIbShenjingVisitorPadGetinfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadGetinfoAPIRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.getinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadGetinfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
