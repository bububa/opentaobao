package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsOfficialaccountOfflineAPIRequest
聚石塔公众号下线 API请求
taobao.jst.sms.officialaccount.offline

聚石塔公众号下线 */
type TaobaoJstSmsOfficialaccountOfflineAPIRequest struct {
	model.Params
	// 公众号下线请求
	_officialAccountOffline *JstBaseRequest
}

// New
