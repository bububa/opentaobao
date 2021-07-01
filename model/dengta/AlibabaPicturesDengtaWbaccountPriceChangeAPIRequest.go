package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest
微博公众号价格变化通知 API请求
alibaba.pictures.dengta.wbaccount.price.change

微博公众号推广价格变更通知接口 */
type AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest struct {
	model.Params
	// 账号id
	_accountId string
	// 转发价格
	_transferPrice string
	// 日期
	_changeTime string
	// 原发价
	_originPrice string
	// id
	_id int64
	// 转发价格 折后价
	_transferPriceAli string
	// 原发价 折后价
	_originPriceAli string
}

// New
