package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWangwangEserviceChatrelationGetAPIRequest
聊天关系获取接口 API请求
taobao.wangwang.eservice.chatrelation.get

获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
2016-09-01， B
2016-09-02， B */
type TaobaoWangwangEserviceChatrelationGetAPIRequest struct {
	model.Params
	// 请求参数
	_chatRelationRequest *ChatRelationRequest
}

// New
