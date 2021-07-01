package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIworkCoreHrsGetpersonAPIRequest
获取神鲸用户基本信息 API请求
alibaba.iwork.core.hrs.getperson

神鲸用户的基本信息查询，根据PERSON_ID或者用户ACCOUNT_ID查询 */
type AlibabaIworkCoreHrsGetpersonAPIRequest struct {
	model.Params
	// 用户ACCOUNT_ID
	_accountId string
	// 用户ID
	_personId int64
	// 应用ID
	_appId string
	// 操作人ID
	_operatorId string
}

// New
