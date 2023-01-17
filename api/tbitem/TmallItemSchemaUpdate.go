package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemSchemaUpdate 天猫根据规则编辑商品
// tmall.item.schema.update
//
// 天猫根据规则编辑商品
func TmallItemSchemaUpdate(clt *core.SDKClient, req *tbitem.TmallItemSchemaUpdateAPIRequest, session string) (*tbitem.TmallItemSchemaUpdateAPIResponse, error) {
	var resp tbitem.TmallItemSchemaUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
