package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkShopGetAPIRequest
淘宝客-推广者-店铺搜索 API请求
taobao.tbk.shop.get

淘宝客店铺查询 */
type TaobaoTbkShopGetAPIRequest struct {
	model.Params
	// 需返回的字段列表
	_fields string
	// 查询词
	_q string
	// 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
	_sort string
	// 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
	_isTmall bool
	// 信用等级下限，1~20
	_startCredit int64
	// 信用等级上限，1~20
	_endCredit int64
	// 淘客佣金比率下限，1~10000
	_startCommissionRate int64
	// 淘客佣金比率上限，1~10000
	_endCommissionRate int64
	// 店铺商品总数下限
	_startTotalAction int64
	// 店铺商品总数上限
	_endTotalAction int64
	// 累计推广商品下限
	_startAuctionCount int64
	// 累计推广商品上限
	_endAuctionCount int64
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 第几页，默认1，1~100
	_pageNo int64
	// 页大小，默认20，1~100
	_pageSize int64
}

// New
