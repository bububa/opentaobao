package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallProductAddSchemaGet 产品发布规则获取接口
// tmall.product.add.schema.get
//
// 获取用户发布产品的规则
func TmallProductAddSchemaGet(clt *core.SDKClient, req *tbitem.TmallProductAddSchemaGetAPIRequest, resp *tbitem.TmallProductAddSchemaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
