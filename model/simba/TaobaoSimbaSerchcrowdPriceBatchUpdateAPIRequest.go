package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest
单品推广搜索人群修改溢价 API请求
taobao.simba.serchcrowd.price.batch.update

单品推广搜索人群修改溢价, 不支持跨推广单元修改 */
type TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 子帐号nick
	_subNick string
	// 需要修改出价的人群包id,批量传入的时候用,分割
	_adgroupCrowdIds []int64
	// 推广单元id
	_adgroupId int64
	// 人群溢价比例，溢价范围[5,300]
	_discount int64
}

// New
