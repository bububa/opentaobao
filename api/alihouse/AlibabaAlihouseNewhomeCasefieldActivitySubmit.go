package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCasefieldActivitySubmit 案场活动维护
// alibaba.alihouse.newhome.casefield.activity.submit
//
// 案场活动维护
func AlibabaAlihouseNewhomeCasefieldActivitySubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
