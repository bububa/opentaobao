package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest
记录用户每日消耗卡路里总量 API请求
alibaba.alihealth.alipaypfm.consume.record

记录用户每日消耗卡路里总量 */
type AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest struct {
	model.Params
	// 用户健康ID
	_userId int64
	// 用户消耗卡路里总量
	_energy int64
	// 记录日期, 格式: yyyy-MM-dd
	_date string
}

// New
