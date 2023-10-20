package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemschemaupdate 天猫根据规则编辑商品
// tmall.item.schema.update
//
// 天猫根据规则编辑商品
func Tmallitemschemaupdate(clt *core.SDKClient, req *tbitem.TmallitemschemaupdateAPIRequest, session string) (*tbitem.TmallitemschemaupdateAPIResponse, error) {
	var resp tbitem.TmallitemschemaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
