package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrinventoryupdate 门店业务同步库存
// tmall.nr.inventory.update
//
// 用于商家每日同步更新门店库存
func Tmallnrinventoryupdate(clt *core.SDKClient, req *tmallnr.TmallnrinventoryupdateAPIRequest, session string) (*tmallnr.TmallnrinventoryupdateAPIResponse, error) {
	var resp tmallnr.TmallnrinventoryupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
