package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

/* TmallNrInventoryUpdate
门店业务同步库存
tmall.nr.inventory.update

用于商家每日同步更新门店库存 */
func TmallNrInventoryUpdate(clt *core.SDKClient, req *tmallnr.TmallNrInventoryUpdateAPIRequest, session string) (*tmallnr.TmallNrInventoryUpdateAPIResponse, error) {
	var resp tmallnr.TmallNrInventoryUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
