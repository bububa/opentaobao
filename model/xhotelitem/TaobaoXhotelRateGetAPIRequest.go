package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRateGetAPIRequest
酒店产品库rate查询 API请求
taobao.xhotel.rate.get

酒店产品库rate查询 */
type TaobaoXhotelRateGetAPIRequest struct {
	model.Params
	// gid酒店商品id
	_gid int64
	// 酒店RPID
	_rpid int64
	// 卖家房型ID, 这是卖家自己系统中的房型ID 注意：需要按照规则组合
	_outRid string
	// 卖家自己系统的Code，简称RateCode
	_rateplanCode string
	// 用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。
	_vendor string
	// RateID
	_rateId int64
}

// New
