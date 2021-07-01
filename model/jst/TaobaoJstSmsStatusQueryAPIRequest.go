package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsStatusQueryAPIRequest
聚石塔公众号状态查询 API请求
taobao.jst.sms.status.query

聚石塔公众号状态查询 */
type TaobaoJstSmsStatusQueryAPIRequest struct {
	model.Params
	// 公众号状态信息查询请求
	_officialAccountStatusQueryRequest *JstBaseRequest
}

// New
