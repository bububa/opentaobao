package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripEmployeeQueryAPIRequest
企业员工查询 API请求
alitrip.btrip.employee.query

企业员工查询 */
type AlitripBtripEmployeeQueryAPIRequest struct {
	model.Params
	// 入参对象。
	_paramOpenEmployeeQueryRequest *OpenEmployeeQueryRequest
}

// New
