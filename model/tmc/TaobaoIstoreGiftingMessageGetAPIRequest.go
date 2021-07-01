package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoIstoreGiftingMessageGetAPIRequest
gifting消息获取 API请求
taobao.istore.gifting.message.get

该api通过参数查询对应的gifting消息 */
type TaobaoIstoreGiftingMessageGetAPIRequest struct {
	model.Params
	// 消息查询条件
	_giftMessageBizCondition *GiftMessageBizCondition
}

// New
