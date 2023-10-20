package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkwmspickmedicinechecksell 联营商药品柜核销
// wdk.wms.pick.medicine.checksell
//
// 联营商药品柜核销
func Wdkwmspickmedicinechecksell(clt *core.SDKClient, req *wdk.WdkwmspickmedicinechecksellAPIRequest, session string) (*wdk.WdkwmspickmedicinechecksellAPIResponse, error) {
	var resp wdk.WdkwmspickmedicinechecksellAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
