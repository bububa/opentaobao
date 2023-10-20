package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// Taobaousergrowthdhhdeliverybatchask 广告曝光前判定批量接口V2
// taobao.usergrowth.dhh.delivery.batchask
//
// 广告曝光前判定批量接口V2
func Taobaousergrowthdhhdeliverybatchask(clt *core.SDKClient, req *usergrowth.TaobaousergrowthdhhdeliverybatchaskAPIRequest, session string) (*usergrowth.TaobaousergrowthdhhdeliverybatchaskAPIResponse, error) {
	var resp usergrowth.TaobaousergrowthdhhdeliverybatchaskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
