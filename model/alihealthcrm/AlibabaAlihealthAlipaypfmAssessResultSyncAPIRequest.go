package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest
用户测评结果回传接口 API请求
alibaba.alihealth.alipaypfm.assess.result.sync

用户测评结果回传接口 */
type AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest struct {
	model.Params
	// userId
	_userId int64
	// 测评类型
	_assessType string
	// 测评结果
	_assessResult string
	// 测评结果冗余字段
	_refrenceResult string
}

// New
