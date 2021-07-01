package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelIntlAriNotifyAPIRequest
国际酒店集团价库变更通知 API请求
taobao.xhotel.intl.ari.notify

国际酒店集团价库变更时通知变更内容，平台及时更新价库信息，保证价库新鲜度 */
type TaobaoXhotelIntlAriNotifyAPIRequest struct {
	model.Params
	// 缓存变更
	_cacheChangeList []CacheChangeInfo
}

// New
