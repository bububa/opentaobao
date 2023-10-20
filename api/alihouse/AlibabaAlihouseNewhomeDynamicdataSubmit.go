package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeDynamicdataSubmit 提交小区动态信息
// alibaba.alihouse.newhome.dynamicdata.submit
//
// 提交小区动态信息
func AlibabaAlihouseNewhomeDynamicdataSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
