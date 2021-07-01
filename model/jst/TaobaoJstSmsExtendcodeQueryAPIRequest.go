package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsExtendcodeQueryAPIRequest
聚石塔扩展码查询 API请求
taobao.jst.sms.extendcode.query

聚石塔扩展码查询 */
type TaobaoJstSmsExtendcodeQueryAPIRequest struct {
	model.Params
	// 扩展码查询请求
	_extendCodeQueryRequest *ExtendCodeQueryRequest
}

// New
