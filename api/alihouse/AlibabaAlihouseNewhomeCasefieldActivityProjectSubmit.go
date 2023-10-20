package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCasefieldActivityProjectSubmit 案场活动楼盘维护
// alibaba.alihouse.newhome.casefield.activity.project.submit
//
// 案场活动楼盘维护
func AlibabaAlihouseNewhomeCasefieldActivityProjectSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
