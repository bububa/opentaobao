package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemcombineget 组合商品获取接口
// tmall.item.combine.get
//
// 查询组合商品的SKU信息
func Tmallitemcombineget(clt *core.SDKClient, req *tbitem.TmallitemcombinegetAPIRequest, session string) (*tbitem.TmallitemcombinegetAPIResponse, error) {
	var resp tbitem.TmallitemcombinegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
