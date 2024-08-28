package alihealthmedical

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

// AlibabaAlihealthMedicalOrderQuery 三方机构查询订单详情接口
// alibaba.alihealth.medical.order.query
//
// 查询订单详情，包括评价
func AlibabaAlihealthMedicalOrderQuery(ctx context.Context, clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalOrderQueryAPIRequest, resp *alihealthmedical.AlibabaAlihealthMedicalOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
