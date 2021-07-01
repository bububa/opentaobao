package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest
点击标签后排序接口 API请求
alibaba.alihealth.pregnancy.taginfo.addrank

备孕管理--点击标签后排序接口 */
type AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
	// 标签编码，例如备孕标签为5122
	_tagCode string
}

// New
