package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectEcode 楼盘e码更新
// alibaba.alihouse.newhome.project.ecode
//
// 更新楼盘ecode
func AlibabaAlihouseNewhomeProjectEcode(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectEcodeAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectEcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
