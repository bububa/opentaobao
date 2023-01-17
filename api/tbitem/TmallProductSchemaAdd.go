package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallProductSchemaAdd 使用Schema文件发布一个产品
// tmall.product.schema.add
//
// Schema体系发布一个产品
func TmallProductSchemaAdd(clt *core.SDKClient, req *tbitem.TmallProductSchemaAddAPIRequest, session string) (*tbitem.TmallProductSchemaAddAPIResponse, error) {
	var resp tbitem.TmallProductSchemaAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
