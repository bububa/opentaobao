package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardSearchcardAPIRequest
搜索卡实例列表(支持号段查询) API请求
alibaba.alsc.crm.card.searchcard

搜索卡实例列表(支持号段查询) */
type AlibabaAlscCrmCardSearchcardAPIRequest struct {
	model.Params
	// 请求对象
	_paramSearchCardOpenReq *SearchCardOpenReq
}

// New
