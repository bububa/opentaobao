package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRightBindBack 权限回流
// alibaba.alihouse.newhome.right.bind.back
//
// 权限回流
func AlibabaAlihouseNewhomeRightBindBack(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRightBindBackAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeRightBindBackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
