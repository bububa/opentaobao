package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest
获取用户上架在线销售的全部宝贝 API请求
taobao.simba.adgroup.onlineitemsvon.get

获取用户上架在线销售的全部宝贝 */
type TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 排序字段，starts：按开始时间排序bidCount:按销量排序
	_orderField string
	// 排序，true:降序， false:升序
	_orderBy bool
	// 页尺寸，最大200
	_pageSize int64
	// 页码，从1开始,最大50。最大只能获取1W个宝贝
	_pageNo int64
	// 推广单元类型 101001005代表标准推广，101001014代表销量明星推广
	_productId int64
}

// New
