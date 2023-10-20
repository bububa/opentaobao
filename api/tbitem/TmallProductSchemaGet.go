package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallProductSchemaGet 产品信息获取schema获取
// tmall.product.schema.get
//
// 产品信息获取接口schema形式返回
func TmallProductSchemaGet(clt *core.SDKClient, req *tbitem.TmallProductSchemaGetAPIRequest, resp *tbitem.TmallProductSchemaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
