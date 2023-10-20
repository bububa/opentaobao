package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuCategorySchemaLevelGet (新)层级属性获取
// alibaba.icbu.category.schema.level.get
//
// 将表单中层级属性的子属性返回
func AlibabaIcbuCategorySchemaLevelGet(clt *core.SDKClient, req *icbu.AlibabaIcbuCategorySchemaLevelGetAPIRequest, resp *icbu.AlibabaIcbuCategorySchemaLevelGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
