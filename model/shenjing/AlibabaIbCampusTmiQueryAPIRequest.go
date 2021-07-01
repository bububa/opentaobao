package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIbCampusTmiQueryAPIRequest
IB智慧园区-查询TMI流水 API请求
alibaba.ib.campus.tmi.query

获取特定银行账户的银行流水 */
type AlibabaIbCampusTmiQueryAPIRequest struct {
	model.Params
	// 查询参数
	_accountQueryReqDto *AccountQueryReqDto
}

// New
