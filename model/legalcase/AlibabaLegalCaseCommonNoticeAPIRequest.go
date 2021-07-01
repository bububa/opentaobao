package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalCaseCommonNoticeAPIRequest
消息通知 API请求
alibaba.legal.case.common.notice

同步通知给诉讼系统 */
type AlibabaLegalCaseCommonNoticeAPIRequest struct {
	model.Params
	// 案件id
	_caseId int64
	// 委托id
	_entrustId int64
	// adminicular_evidence（补充证据）risk_early_warning（风险预警）
	_type string
	// 消息模型
	_noticeModel *NoticeModel
}

// New
