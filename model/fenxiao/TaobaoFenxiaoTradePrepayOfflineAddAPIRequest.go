package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoTradePrepayOfflineAddAPIRequest
线下预存款流水增加 API请求
taobao.fenxiao.trade.prepay.offline.add

渠道分销供应商上传线下流水预存款（增加） */
type TaobaoFenxiaoTradePrepayOfflineAddAPIRequest struct {
	model.Params
	// 增加流水
	_offlineAddPrepayParam *TopOfflineAddPrepayDto
}

// New
