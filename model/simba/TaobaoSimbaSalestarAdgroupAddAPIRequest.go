package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarAdgroupAddAPIRequest
(新)创建一个推广组 API请求
taobao.simba.salestar.adgroup.add

创建一个推广组 */
type TaobaoSimbaSalestarAdgroupAddAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
	// 商品Id
	_itemId int64
	// 创意标题，最多20个汉字
	_title string
	// 创意图片地址，必须是商品的图片之一
	_imgUrl string
}

// New
