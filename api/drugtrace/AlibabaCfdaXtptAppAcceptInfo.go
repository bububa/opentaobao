package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaCfdaXtptAppAcceptInfo 协同平台数据下行接口
// alibaba.cfda.xtpt.app.accept.info
//
// 协同平台数据下行接口
func AlibabaCfdaXtptAppAcceptInfo(clt *core.SDKClient, req *drugtrace.AlibabaCfdaXtptAppAcceptInfoAPIRequest, resp *drugtrace.AlibabaCfdaXtptAppAcceptInfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
