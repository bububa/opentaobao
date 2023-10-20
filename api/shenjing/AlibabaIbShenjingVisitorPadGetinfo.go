package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// AlibabaIbShenjingVisitorPadGetinfo 获取OSS上传参数
// alibaba.ib.shenjing.visitor.pad.getinfo
//
// PAD 端获取OSS上传参数，向OSS服务器上传图片。
func AlibabaIbShenjingVisitorPadGetinfo(clt *core.SDKClient, req *shenjing.AlibabaIbShenjingVisitorPadGetinfoAPIRequest, resp *shenjing.AlibabaIbShenjingVisitorPadGetinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
