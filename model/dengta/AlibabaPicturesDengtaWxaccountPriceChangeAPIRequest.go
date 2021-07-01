package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest
微信公众号价格变化通知 API请求
alibaba.pictures.dengta.wxaccount.price.change

微信公众号推广价格变更通知接口 */
type AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest struct {
	model.Params
	// id
	_accountId string
	// 变更时间
	_changeTime string
	// 单图文
	_single string
	// 多图文第三条及以后
	_other string
	// 多图文第二条
	_second string
	// 多图文第一条
	_first string
	// 账号id
	_id int64
	// 多图文第一条 折后价
	_firstAli string
	// 多图文第二条 折后价
	_secondAli string
	// 单图文 折后价
	_singleAli string
	// 多图文第三条及以后 折后价
	_otherAli string
}

// New
