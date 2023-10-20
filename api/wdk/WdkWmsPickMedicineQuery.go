package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkWmsPickMedicineQuery 查询拣货单中的药品信息
// wdk.wms.pick.medicine.query
//
// 联营商药机查询拣货单中的药品信息
func WdkWmsPickMedicineQuery(clt *core.SDKClient, req *wdk.WdkWmsPickMedicineQueryAPIRequest, resp *wdk.WdkWmsPickMedicineQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
