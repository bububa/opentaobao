package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Alibabaitemeditschemaget 商品编辑获取schema信息
// alibaba.item.edit.schema.get
//
// 商品编辑时，获取商品规则信息
func Alibabaitemeditschemaget(clt *core.SDKClient, req *tbitem.AlibabaitemeditschemagetAPIRequest, session string) (*tbitem.AlibabaitemeditschemagetAPIResponse, error) {
	var resp tbitem.AlibabaitemeditschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
