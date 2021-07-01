package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupAddAPIRequest
创建一个推广组 API请求
taobao.simba.adgroup.add

创建一个推广组 */
type TaobaoSimbaAdgroupAddAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
	// 商品Id
	_itemId int64
	// 推广组默认出价，单位为分，不能小于5 不能大于日最高限额
	_defaultPrice int64
	// 创意标题，最多20个汉字
	_title string
	// 创意图片地址，必须是商品的图片之一
	_imgUrl string
}

// New
