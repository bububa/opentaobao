package eleenterpriserestaurant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseRestaurantSearchAPIRequest
餐厅列表 API请求
alibaba.ele.enterprise.restaurant.search

餐厅列表 */
type AlibabaEleEnterpriseRestaurantSearchAPIRequest struct {
	model.Params
	// longitude和latitude用英文逗号分隔
	_geo string
	// 首次查询无需传入，后续需要传入前次返回
	_rankId string
	// 查询起始位置，默认为0。如果传的是10，那么餐厅会从第11个开始返回
	_start int64
	// 查询数量，默认是10，最大50
	_limit int64
	// 人均消费金额上限，需要高于costFrom，不传表示不限
	_costTo int64
	// 人均消费金额下限，最低为0，不传表示不限
	_costFrom int64
	// 是否支持食安保（0-不限，1-支持食安保）不传表示不限
	_insurance int64
	// 是否可开发票（0-不限，1-可开发票）不传表示不限
	_invoice int64
	// 是否品牌商家（0-不限，1-品牌商家）不传表示不限
	_isPremium int64
	// 是否新店（0-不限，1-新店）不传表示不限
	_newRestaurant int64
	// 配送方式（0-不限， 1-蜂鸟专送）不传表示不限
	_deliveryMode int64
	// 排序选项（1-默认排序（热门）， 2-评价星级由高到低， 3-起送价由低到高， 4-销量由高到低， 5-送餐速度由快到慢， 6-餐厅距离由近到远， 7-订单量由高到低）
	_orderBy int64
	// 餐厅分类ids
	_categoryIds []int64
	// 是否筛选支持预定 0:不需要 1:需要（不传该字段则不筛选）
	_isBookable string
	// 是否支持跨天预定 1:需要（不传该字段则不筛选）
	_crossDayBooking string
}

// New
