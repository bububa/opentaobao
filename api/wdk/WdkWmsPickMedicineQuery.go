package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkWmsPickMedicineQuery 查询拣货单中的药品信息
// wdk.wms.pick.medicine.query
//
// 联营商药机查询拣货单中的药品信息
func WdkWmsPickMedicineQuery(ctx context.Context, clt *core.SDKClient, req *wdk.WdkWmsPickMedicineQueryAPIRequest, resp *wdk.WdkWmsPickMedicineQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
