package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcMessagesProduceAPIRequest
批量发送消息 API请求
taobao.tmc.messages.produce

批量发送消息 */
type TaobaoTmcMessagesProduceAPIRequest struct {
	model.Params
	// tmc消息列表, 最多50条，元素结构与taobao.tmc.message.produce一致，用json表示的消息列表。例如：[{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"},{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"}]
	_messages []TmcPublishMessage
}

// New
