package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// Taobaousergrowthdhhdeliveryask 广告曝光前判定接口V2
// taobao.usergrowth.dhh.delivery.ask
//
// 提供给媒体在曝光广告前调用
func Taobaousergrowthdhhdeliveryask(clt *core.SDKClient, req *usergrowth.TaobaousergrowthdhhdeliveryaskAPIRequest, session string) (*usergrowth.TaobaousergrowthdhhdeliveryaskAPIResponse, error) {
	var resp usergrowth.TaobaousergrowthdhhdeliveryaskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
