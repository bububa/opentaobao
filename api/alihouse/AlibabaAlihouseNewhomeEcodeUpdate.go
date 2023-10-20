package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeEcodeUpdate 新房货变更E码
// alibaba.alihouse.newhome.ecode.update
//
// 新房货变更E码
func AlibabaAlihouseNewhomeEcodeUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeEcodeUpdateAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeEcodeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
