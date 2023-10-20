package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemUpdateSchemaGet 天猫编辑商品规则获取
// tmall.item.update.schema.get
//
// Schema方式编辑天猫商品时，编辑商品规则获取
func TmallItemUpdateSchemaGet(clt *core.SDKClient, req *tbitem.TmallItemUpdateSchemaGetAPIRequest, resp *tbitem.TmallItemUpdateSchemaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
