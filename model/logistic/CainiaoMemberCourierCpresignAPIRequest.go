package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoMemberCourierCpresignAPIRequest
cp清理离职用户信息 API请求
cainiao.member.courier.cpresign

CP清理内部离职的用户信息 */
type CainiaoMemberCourierCpresignAPIRequest struct {
	model.Params
	// 菜鸟用户id
	_accountId int64
}

// New
