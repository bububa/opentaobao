package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemEditSchemaGet 商品编辑获取schema信息
// alibaba.item.edit.schema.get
//
// 商品编辑时，获取商品规则信息
func AlibabaItemEditSchemaGet(clt *core.SDKClient, req *tbitem.AlibabaItemEditSchemaGetAPIRequest, session string) (*tbitem.AlibabaItemEditSchemaGetAPIResponse, error) {
	var resp tbitem.AlibabaItemEditSchemaGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
