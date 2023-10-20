package gameact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

// Taobaodeactivityinfoget 获取活动信息
// taobao.de.activity.info.get
//
// 根据appKey和活动id获取活动
func Taobaodeactivityinfoget(clt *core.SDKClient, req *gameact.TaobaodeactivityinfogetAPIRequest, session string) (*gameact.TaobaodeactivityinfogetAPIResponse, error) {
	var resp gameact.TaobaodeactivityinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
