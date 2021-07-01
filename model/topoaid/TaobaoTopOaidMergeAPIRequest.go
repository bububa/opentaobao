package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopOaidMergeAPIRequest
OAID订单合并 API请求
taobao.top.oaid.merge

基于OAID（收件人ID， Open Addressee ID)做订单合并，确保相同收件人信息的订单合并到相同组。 */
type TaobaoTopOaidMergeAPIRequest struct {
	model.Params
	// 合单请求列表，最多支持100个。
	_mergeList []OrderMerge
}

// New
