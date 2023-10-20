package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectSalestime 楼盘销售时刻同步
// alibaba.alihouse.newhome.project.salestime
//
// 楼盘销售时刻同步
func AlibabaAlihouseNewhomeProjectSalestime(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectSalestimeAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectSalestimeAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectSalestimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
