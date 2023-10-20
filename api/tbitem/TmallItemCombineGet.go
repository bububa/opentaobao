package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemCombineGet 组合商品获取接口
// tmall.item.combine.get
//
// 查询组合商品的SKU信息
func TmallItemCombineGet(clt *core.SDKClient, req *tbitem.TmallItemCombineGetAPIRequest, resp *tbitem.TmallItemCombineGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
