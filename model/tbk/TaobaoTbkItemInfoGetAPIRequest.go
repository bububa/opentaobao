package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkItemInfoGetAPIRequest
淘宝客-公用-淘宝客商品详情查询(简版) API请求
taobao.tbk.item.info.get

淘宝客商品详情查询（简版） */
type TaobaoTbkItemInfoGetAPIRequest struct {
	model.Params
	// 商品ID串，用,分割，最大40个
	_numIids string
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	_ip string
}

// New
