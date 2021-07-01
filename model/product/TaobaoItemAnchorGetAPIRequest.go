package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemAnchorGetAPIRequest
获取可用宝贝描述规范化模块 API请求
taobao.item.anchor.get

根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点 */
type TaobaoItemAnchorGetAPIRequest struct {
	model.Params
	// 宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1.(最小值：-1，最大值：1)
	_type int64
	// 对应类目编号
	_catId int64
}

// New
