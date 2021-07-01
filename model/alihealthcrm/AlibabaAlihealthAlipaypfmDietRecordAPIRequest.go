package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlipaypfmDietRecordAPIRequest
用户每日摄入卡路里总量回传接口 API请求
alibaba.alihealth.alipaypfm.diet.record

用户每日摄入卡路里总量回传接口 */
type AlibabaAlihealthAlipaypfmDietRecordAPIRequest struct {
	model.Params
	// 用户健康ID
	_userId int64
	// 记录日期，format：yyyy-MM-dd
	_date string
	// 累积摄入卡路里
	_energy int64
}

// New
