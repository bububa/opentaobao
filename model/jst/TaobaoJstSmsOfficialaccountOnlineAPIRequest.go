package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsOfficialaccountOnlineAPIRequest
聚石塔公众号上线 API请求
taobao.jst.sms.officialaccount.online

聚石塔公众号上线 */
type TaobaoJstSmsOfficialaccountOnlineAPIRequest struct {
	model.Params
	// 公众号上线请求参数
	_officialAccountOnlineRequest *JstBaseRequest
}

// New
