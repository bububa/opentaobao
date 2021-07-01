package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest
获取全国疫情统计数据 API请求
alibaba.alihealth.nocov.alldiseaseinfo.get

获取全国疫情统计数据 */
type AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest struct {
	model.Params
	// 省的
	_province string
	// 城市
	_city string
	// 城市code
	_cityCode string
}

// New
