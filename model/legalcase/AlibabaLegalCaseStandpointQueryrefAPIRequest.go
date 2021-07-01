package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalCaseStandpointQueryrefAPIRequest
查询推送口径信息 API请求
alibaba.legal.case.standpoint.queryref

查询推送口径信息 */
type AlibabaLegalCaseStandpointQueryrefAPIRequest struct {
	model.Params
	// 案件id
	_suitId int64
	// 委托id
	_entrustId int64
	// 查询的口径id列表
	_standpointIds []int64
}

// New
