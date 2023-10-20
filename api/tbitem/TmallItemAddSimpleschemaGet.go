package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemaddsimpleschemaget 天猫发布商品规则获取
// tmall.item.add.simpleschema.get
//
// 通过商家信息获取商品发布字段和规则。
func Tmallitemaddsimpleschemaget(clt *core.SDKClient, req *tbitem.TmallitemaddsimpleschemagetAPIRequest, session string) (*tbitem.TmallitemaddsimpleschemagetAPIResponse, error) {
	var resp tbitem.TmallitemaddsimpleschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
