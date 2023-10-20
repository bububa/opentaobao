package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemupdateschemaget 天猫编辑商品规则获取
// tmall.item.update.schema.get
//
// Schema方式编辑天猫商品时，编辑商品规则获取
func Tmallitemupdateschemaget(clt *core.SDKClient, req *tbitem.TmallitemupdateschemagetAPIRequest, session string) (*tbitem.TmallitemupdateschemagetAPIResponse, error) {
	var resp tbitem.TmallitemupdateschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
