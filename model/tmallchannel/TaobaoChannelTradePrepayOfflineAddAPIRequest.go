package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoChannelTradePrepayOfflineAddAPIRequest
渠道分销供应商上传线下流水预存款（增加） API请求
taobao.channel.trade.prepay.offline.add

渠道分销供应商上传线下流水预存款（增加） */
type TaobaoChannelTradePrepayOfflineAddAPIRequest struct {
	model.Params
	// 增加流水
	_offlineAddPrepayParam *TopOfflineAddPrepayDto
}

// New
