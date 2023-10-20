package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemtemplatesget 获取用户宝贝详情页模板名称
// taobao.item.templates.get
//
// 查询当前登录用户的店铺的宝贝详情页的模板名称
func Taobaoitemtemplatesget(clt *core.SDKClient, req *tbitem.TaobaoitemtemplatesgetAPIRequest, session string) (*tbitem.TaobaoitemtemplatesgetAPIResponse, error) {
	var resp tbitem.TaobaoitemtemplatesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
