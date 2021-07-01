package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKaolaScitemAddAPIRequest
考拉货品新增接口 API请求
taobao.kaola.scitem.add

考拉货品新增接口 */
type TaobaoKaolaScitemAddAPIRequest struct {
	model.Params
	// 待新增的货品
	_cnsku *CnskuDto
	// 新增选项
	_option *AddCnskuOption
}

// New
