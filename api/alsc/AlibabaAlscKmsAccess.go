package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscKmsAccess 本地生活风控数据接入
// alibaba.alsc.kms.access
//
// 第三方使用本地生活数据对外提供服务，上报访问日志信息接口
func AlibabaAlscKmsAccess(clt *core.SDKClient, req *alsc.AlibabaAlscKmsAccessAPIRequest, resp *alsc.AlibabaAlscKmsAccessAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
