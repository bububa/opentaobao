package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemHscodeDetailGet 通过hscode获取计量单位
// tmall.item.hscode.detail.get
//
// 通过hscode获取计量单位和销售单位
func TmallItemHscodeDetailGet(clt *core.SDKClient, req *tbitem.TmallItemHscodeDetailGetAPIRequest, session string) (*tbitem.TmallItemHscodeDetailGetAPIResponse, error) {
	var resp tbitem.TmallItemHscodeDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
