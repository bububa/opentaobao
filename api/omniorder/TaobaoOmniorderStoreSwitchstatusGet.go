package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstoreswitchstatusget switchstatus.get
// taobao.omniorder.store.switchstatus.get
//
// 查询门店发货、门店自提状态
func Taobaoomniorderstoreswitchstatusget(clt *core.SDKClient, req *omniorder.TaobaoomniorderstoreswitchstatusgetAPIRequest, session string) (*omniorder.TaobaoomniorderstoreswitchstatusgetAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstoreswitchstatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
