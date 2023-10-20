package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// Taobaousergrowthtaskconfigget 用户增长营销玩法配置查询
// taobao.usergrowth.task.config.get
//
// 用户增长营销玩法配置查询
func Taobaousergrowthtaskconfigget(clt *core.SDKClient, req *usergrowth.TaobaousergrowthtaskconfiggetAPIRequest, session string) (*usergrowth.TaobaousergrowthtaskconfiggetAPIResponse, error) {
	var resp usergrowth.TaobaousergrowthtaskconfiggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
