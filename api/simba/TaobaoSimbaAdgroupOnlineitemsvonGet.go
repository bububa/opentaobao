package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupOnlineitemsvonGet 获取用户上架在线销售的全部宝贝
// taobao.simba.adgroup.onlineitemsvon.get
//
// 获取用户上架在线销售的全部宝贝
func TaobaoSimbaAdgroupOnlineitemsvonGet(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest, session string) (*simba.TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse, error) {
	var resp simba.TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
