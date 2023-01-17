package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemCombineGet 组合商品获取接口
// tmall.item.combine.get
//
// 查询组合商品的SKU信息
func TmallItemCombineGet(clt *core.SDKClient, req *tbitem.TmallItemCombineGetAPIRequest, session string) (*tbitem.TmallItemCombineGetAPIResponse, error) {
	var resp tbitem.TmallItemCombineGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
