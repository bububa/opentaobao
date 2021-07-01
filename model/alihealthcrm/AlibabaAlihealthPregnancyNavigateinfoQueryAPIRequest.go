package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest
查询底导数据 API请求
alibaba.alihealth.pregnancy.navigateinfo.query

备孕管理--获取底部导航信息 */
type AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
}

// New
