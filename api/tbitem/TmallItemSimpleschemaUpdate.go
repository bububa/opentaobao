package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemsimpleschemaupdate 天猫简化编辑商品
// tmall.item.simpleschema.update
//
// 国外大商家天猫简化编辑商品
func Tmallitemsimpleschemaupdate(clt *core.SDKClient, req *tbitem.TmallitemsimpleschemaupdateAPIRequest, session string) (*tbitem.TmallitemsimpleschemaupdateAPIResponse, error) {
	var resp tbitem.TmallitemsimpleschemaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
