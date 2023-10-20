package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaowlbstoresbaseinfoget 商家按照仓的类型获取仓库接口
// taobao.wlb.stores.baseinfo.get
//
// 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
func Taobaowlbstoresbaseinfoget(clt *core.SDKClient, req *tblogistics.TaobaowlbstoresbaseinfogetAPIRequest, session string) (*tblogistics.TaobaowlbstoresbaseinfogetAPIResponse, error) {
	var resp tblogistics.TaobaowlbstoresbaseinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
