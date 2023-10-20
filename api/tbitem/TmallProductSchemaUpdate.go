package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallProductSchemaUpdate 产品更新接口
// tmall.product.schema.update
//
// 产品更新接口
func TmallProductSchemaUpdate(clt *core.SDKClient, req *tbitem.TmallProductSchemaUpdateAPIRequest, resp *tbitem.TmallProductSchemaUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
