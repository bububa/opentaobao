package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardQueryTemplateAPIRequest
查询卡模板详情 API请求
alibaba.alsc.crm.card.query.template

查询卡模板详情 */
type AlibabaAlscCrmCardQueryTemplateAPIRequest struct {
	model.Params
	// 请求对象
	_paramQueryCardTemplateOpenReq *QueryCardTemplateOpenReq
}

// New
