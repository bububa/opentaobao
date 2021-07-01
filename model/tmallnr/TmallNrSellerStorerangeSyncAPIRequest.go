package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrSellerStorerangeSyncAPIRequest
同步商户中心服务范围 API请求
tmall.nr.seller.storerange.sync

同步商户中心服务范围 */
type TmallNrSellerStorerangeSyncAPIRequest struct {
	model.Params
	// 业务身份标识,dss定时送；self_day 自配日达；self_hour 自配小时达
	_bizIdentity string
	// 系统自动生成
	_reqDTOList []SyncServiceRangeRequestDto
	// 卖家id，有可能和登录seller不是同一个id
	_sellerId int64
}

// New
