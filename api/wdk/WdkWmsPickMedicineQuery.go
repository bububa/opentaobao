package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkwmspickmedicinequery 查询拣货单中的药品信息
// wdk.wms.pick.medicine.query
//
// 联营商药机查询拣货单中的药品信息
func Wdkwmspickmedicinequery(clt *core.SDKClient, req *wdk.WdkwmspickmedicinequeryAPIRequest, session string) (*wdk.WdkwmspickmedicinequeryAPIResponse, error) {
	var resp wdk.WdkwmspickmedicinequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
