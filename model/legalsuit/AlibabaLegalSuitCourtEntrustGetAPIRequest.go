package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitCourtEntrustGetAPIRequest
委托开庭服务查询 API请求
alibaba.legal.suit.court.entrust.get

查询委托开庭信息 */
type AlibabaLegalSuitCourtEntrustGetAPIRequest struct {
	model.Params
	// 委托ID
	_entrustId int64
	// 案件ID
	_suitId int64
}

// New
