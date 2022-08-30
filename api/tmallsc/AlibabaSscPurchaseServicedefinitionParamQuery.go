package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaSscPurchaseServicedefinitionParamQuery 查询采购服务定义参数信息
// alibaba.ssc.purchase.servicedefinition.param.query
//
// 查询采购服务定义参数信息
func AlibabaSscPurchaseServicedefinitionParamQuery(clt *core.SDKClient, req *tmallsc.AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest, session string) (*tmallsc.AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse, error) {
	var resp tmallsc.AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
