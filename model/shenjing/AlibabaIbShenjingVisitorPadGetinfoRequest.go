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
type AlibabaIbShenjingVisitorPadGetinfoRequest struct {
    model.Params
}

// 初始化AlibabaIbShenjingVisitorPadGetinfoRequest对象
func NewAlibabaIbShenjingVisitorPadGetinfoRequest() *AlibabaIbShenjingVisitorPadGetinfoRequest{
    return &AlibabaIbShenjingVisitorPadGetinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadGetinfoRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.getinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadGetinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
