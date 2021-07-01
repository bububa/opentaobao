package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbStoresBaseinfoGetAPIRequest
商家按照仓的类型获取仓库接口 API请求
taobao.wlb.stores.baseinfo.get

通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库 */
type TaobaoWlbStoresBaseinfoGetAPIRequest struct {
	model.Params
	// 0.商家仓库.1.菜鸟仓库.2全部被划分的仓库
	_type int64
}

// New
