package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitCourtAfterPushAPIRequest
更新或者新增庭后信息 API请求
alibaba.legal.suit.court.after.push

供外部ISV供应商 推送庭后信息给集团诉讼 */
type AlibabaLegalSuitCourtAfterPushAPIRequest struct {
	model.Params
	// 庭后信息
	_afterCourtInfoModel *AfterCourtInfoModel
}

// New
