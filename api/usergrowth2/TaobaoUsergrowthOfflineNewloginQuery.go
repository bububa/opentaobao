package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthOfflineNewloginQuery 线下新登数据实时查询接口
// taobao.usergrowth.offline.newlogin.query
//
// 线下新登数据实时查询接口
func TaobaoUsergrowthOfflineNewloginQuery(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthOfflineNewloginQueryAPIRequest, session string) (*usergrowth2.TaobaoUsergrowthOfflineNewloginQueryAPIResponse, error) {
	var resp usergrowth2.TaobaoUsergrowthOfflineNewloginQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
