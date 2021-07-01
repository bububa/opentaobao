package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoShopUpdateAPIRequest
更新店铺基本信息 API请求
taobao.shop.update

目前只支持标题、公告和描述的更新 */
type TaobaoShopUpdateAPIRequest struct {
	model.Params
	// 店铺标题。不超过30个字符；过滤敏感词，如淘咖啡、阿里巴巴等。title, bulletin和desc至少必须传一个
	_title string
	// 店铺公告。不超过1024个字符
	_bulletin string
	// 店铺描述。10～2000个字符以内
	_desc string
}

// New
