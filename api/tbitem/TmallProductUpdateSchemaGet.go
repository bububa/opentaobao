package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallProductUpdateSchemaGet 产品更新规则获取接口
// tmall.product.update.schema.get
//
// 获取用户更新产品的规则
func TmallProductUpdateSchemaGet(clt *core.SDKClient, req *tbitem.TmallProductUpdateSchemaGetAPIRequest, resp *tbitem.TmallProductUpdateSchemaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
