package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoWlbStoresBaseinfoGet 商家按照仓的类型获取仓库接口
// taobao.wlb.stores.baseinfo.get
//
// 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
func TaobaoWlbStoresBaseinfoGet(clt *core.SDKClient, req *logistic.TaobaoWlbStoresBaseinfoGetAPIRequest, session string) (*logistic.TaobaoWlbStoresBaseinfoGetAPIResponse, error) {
	var resp logistic.TaobaoWlbStoresBaseinfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
