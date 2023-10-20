package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectSalestime 楼盘销售时刻同步
// alibaba.alihouse.newhome.project.salestime
//
// 楼盘销售时刻同步
func AlibabaAlihouseNewhomeProjectSalestime(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectSalestimeAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectSalestimeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
