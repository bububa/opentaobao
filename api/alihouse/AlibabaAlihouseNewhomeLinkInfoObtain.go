package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeLinkInfoObtain 落地页获取
// alibaba.alihouse.newhome.link.info.obtain
//
// 落地页获取
func AlibabaAlihouseNewhomeLinkInfoObtain(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
