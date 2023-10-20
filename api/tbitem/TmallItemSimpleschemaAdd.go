package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemsimpleschemaadd 天猫简化发布商品
// tmall.item.simpleschema.add
//
// 天猫简化版schema发布商品。
func Tmallitemsimpleschemaadd(clt *core.SDKClient, req *tbitem.TmallitemsimpleschemaaddAPIRequest, session string) (*tbitem.TmallitemsimpleschemaaddAPIResponse, error) {
	var resp tbitem.TmallitemsimpleschemaaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
