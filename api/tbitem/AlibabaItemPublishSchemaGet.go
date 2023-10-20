package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemPublishSchemaGet 获取商品发布规则信息
// alibaba.item.publish.schema.get
//
// 新商品发布，获取商品发布规则信息
func AlibabaItemPublishSchemaGet(clt *core.SDKClient, req *tbitem.AlibabaItemPublishSchemaGetAPIRequest, resp *tbitem.AlibabaItemPublishSchemaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
