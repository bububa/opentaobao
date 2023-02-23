package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemTemplatesGet 获取用户宝贝详情页模板名称
// taobao.item.templates.get
//
// 查询当前登录用户的店铺的宝贝详情页的模板名称
func TaobaoItemTemplatesGet(clt *core.SDKClient, req *tbitem.TaobaoItemTemplatesGetAPIRequest, session string) (*tbitem.TaobaoItemTemplatesGetAPIResponse, error) {
	var resp tbitem.TaobaoItemTemplatesGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
