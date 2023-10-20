package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupbigbagcontentcancel 进口大包取消
// cainiao.global.im.pickup.bigbag.content.cancel
//
// 进口大包取消
func Cainiaoglobalimpickupbigbagcontentcancel(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupbigbagcontentcancelAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupbigbagcontentcancelAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupbigbagcontentcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
