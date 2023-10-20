package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabasscpurchaseservicedefinitionparamquery 查询采购服务定义参数信息
// alibaba.ssc.purchase.servicedefinition.param.query
//
// 查询采购服务定义参数信息
func Alibabasscpurchaseservicedefinitionparamquery(clt *core.SDKClient, req *tmallsc.AlibabasscpurchaseservicedefinitionparamqueryAPIRequest, session string) (*tmallsc.AlibabasscpurchaseservicedefinitionparamqueryAPIResponse, error) {
	var resp tmallsc.AlibabasscpurchaseservicedefinitionparamqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
