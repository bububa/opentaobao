package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemHscodeDetailGet 通过hscode获取计量单位
// tmall.item.hscode.detail.get
//
// 通过hscode获取计量单位和销售单位
func TmallItemHscodeDetailGet(clt *core.SDKClient, req *tbitem.TmallItemHscodeDetailGetAPIRequest, resp *tbitem.TmallItemHscodeDetailGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
