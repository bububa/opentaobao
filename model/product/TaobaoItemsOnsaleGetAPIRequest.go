package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemsOnsaleGetAPIRequest
获取当前会话用户出售中的商品列表 API请求
taobao.items.onsale.get

获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤
只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get 获取
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong> */
type TaobaoItemsOnsaleGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Item商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id,sold_quantity ；字段之间用“,”分隔。不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get 获取。
	_fields string
	// 搜索字段。搜索商品的title。
	_q string
	// 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
	_cid int64
	// 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)
	_sellerCids string
	// 页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
	_pageNo int64
	// 是否参与会员折扣。可选值：true，false。默认不过滤该条件
	_hasDiscount bool
	// 是否橱窗推荐。 可选值：true，false。默认不过滤该条件
	_hasShowcase bool
	// 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间)，sold_quantity（商品销量）,;默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
	_orderBy string
	// 商品是否在淘宝显示
	_isTaobao bool
	// 商品是否在外部网店显示
	_isEx bool
	// 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
	_pageSize int64
	// 起始的修改时间
	_startModified string
	// 结束的修改时间
	_endModified string
	// 是否挂接了达尔文标准产品体系
	_isCspu bool
	// 组合商品
	_isCombine bool
	// 商品类型：a-拍卖,b-一口价
	_auctionType string
}

// New
