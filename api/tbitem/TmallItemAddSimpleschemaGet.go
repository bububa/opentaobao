package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemAddSimpleschemaGet 天猫发布商品规则获取
// tmall.item.add.simpleschema.get
//
// 通过商家信息获取商品发布字段和规则。
func TmallItemAddSimpleschemaGet(clt *core.SDKClient, req *tbitem.TmallItemAddSimpleschemaGetAPIRequest, session string) (*tbitem.TmallItemAddSimpleschemaGetAPIResponse, error) {
	var resp tbitem.TmallItemAddSimpleschemaGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
