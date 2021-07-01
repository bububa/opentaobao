package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiRankDetailGetAPIRequest
排行榜详情 API请求
alibaba.xiami.api.rank.detail.get

虾米排行榜详情数据 */
type AlibabaXiamiApiRankDetailGetAPIRequest struct {
	model.Params
	// 榜单ID
	_billboardId int64
	// 调用来源
	_bizCode string
}

// New
