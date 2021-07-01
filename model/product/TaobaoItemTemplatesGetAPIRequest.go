package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemTemplatesGetAPIRequest
获取用户宝贝详情页模板名称 API请求
taobao.item.templates.get

查询当前登录用户的店铺的宝贝详情页的模板名称 */
type TaobaoItemTemplatesGetAPIRequest struct {
	model.Params
}

// New
