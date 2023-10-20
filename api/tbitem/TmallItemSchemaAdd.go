package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemschemaadd 天猫根据规则发布商品
// tmall.item.schema.add
//
// 天猫TopSchema发布商品。
func Tmallitemschemaadd(clt *core.SDKClient, req *tbitem.TmallitemschemaaddAPIRequest, session string) (*tbitem.TmallitemschemaaddAPIResponse, error) {
	var resp tbitem.TmallitemschemaaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
