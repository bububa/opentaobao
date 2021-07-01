package tmallitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemsExtendSearchAPIRequest
搜索天猫商品 API请求
tmall.items.extend.search

提供天猫商品搜索结果，需要调用精选商品，请改为调用：tmall.selected.items.search */
type TmallItemsExtendSearchAPIRequest struct {
	model.Params
	// 表示搜索的关键字，例如搜索query=nike。当输入关键字为中文时，将对他进行URLEncode的UTF-8格式编码，如 耐克，那么q=%E8%80%90%E5%85%8B。
	_q string
	// 前台类目id，支持多选过滤，cat=catid1,catid2
	_cat string
	// 排序类型。类型包括：s: 人气排序p: 价格从低到高;pd: 价格从高到低;d: 月销量从高到低;td: 总销量从高到低;pt: 按发布时间排序.
	_sort string
	// 品牌的id。支持多选过滤，brand=brand1,brand2
	_brand string
	// 宝贝卖家所在地，中文gbk编码
	_loc string
	// 以“属性id：属性值”的形式传入;
	_prop string
	// 是否包邮，-1为包邮
	_postFee int64
	// 在宝贝页面中进行价格筛选的时候，如果填写了最低价格，就会显示该字段。
	_startPrice float64
	// 在宝贝页面中进行价格筛选的时候，如果填写了最高价格，就会显示该字段。
	_endPrice float64
	// 是否货到付款，1为货到付款
	_supportCod int64
	// 是否多倍积分，1为多倍积分
	_manyPoints int64
	// 显示旺旺在线卖家的宝贝时，wwonline=1
	_wwonline int64
	// 过滤vip宝贝时，vip=1
	_vip int64
	// 过滤搭配减价宝贝时，combo=1
	_combo int64
	// 过滤折扣宝贝时，miaosha=1
	_miaosha int64
	// 是否需要spu聚合的开关:1为关闭，不传表示遵循后端聚合逻辑。默认不作spu聚合。
	_nspu int64
	// 商品标签。支持多选过滤,auction_tag=auction_tag1,auction_tag2,不支持天猫精品库8578
	_auctionTag string
	// 可以根据产品Id搜索属于这个spu的商品。
	_spuid int64
	// 可以根据卖家id搜索属于该卖家的商品
	_userId int64
	// 页码。取值范围：大于零的整数；最大值：100；默认值：1，即默认返回第一页数据。
	_pageNo int64
	// 每页条数。取值范围：大于零的整数；最大值：100；默认值：40
	_pageSize int64
	// 后台类目id，category=categoryId
	_category string
}

// New
