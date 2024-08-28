package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkWmsPickMedicineChecksell 联营商药品柜核销
// wdk.wms.pick.medicine.checksell
//
// 联营商药品柜核销
func WdkWmsPickMedicineChecksell(ctx context.Context, clt *core.SDKClient, req *wdk.WdkWmsPickMedicineChecksellAPIRequest, resp *wdk.WdkWmsPickMedicineChecksellAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
