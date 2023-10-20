package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstoreswitchstatusupdate switchstatus.update
// taobao.omniorder.store.switchstatus.update
//
// 变更门店发货、门店自提状态
func Taobaoomniorderstoreswitchstatusupdate(clt *core.SDKClient, req *omniorder.TaobaoomniorderstoreswitchstatusupdateAPIRequest, session string) (*omniorder.TaobaoomniorderstoreswitchstatusupdateAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstoreswitchstatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
