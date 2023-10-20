package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemSchemaAdd 天猫根据规则发布商品
// tmall.item.schema.add
//
// 天猫TopSchema发布商品。
func TmallItemSchemaAdd(clt *core.SDKClient, req *tbitem.TmallItemSchemaAddAPIRequest, resp *tbitem.TmallItemSchemaAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
