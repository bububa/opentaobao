package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkWmsPickMedicineChecksell 联营商药品柜核销
// wdk.wms.pick.medicine.checksell
//
// 联营商药品柜核销
func WdkWmsPickMedicineChecksell(clt *core.SDKClient, req *wdk.WdkWmsPickMedicineChecksellAPIRequest, session string) (*wdk.WdkWmsPickMedicineChecksellAPIResponse, error) {
	var resp wdk.WdkWmsPickMedicineChecksellAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
