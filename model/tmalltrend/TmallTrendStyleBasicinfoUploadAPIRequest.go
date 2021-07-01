package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTrendStyleBasicinfoUploadAPIRequest
3D款式基本信息同步API API请求
tmall.trend.style.basicinfo.upload

3D款式基本信息同步至天猫趋势中心 */
type TmallTrendStyleBasicinfoUploadAPIRequest struct {
	model.Params
	// 款式基本信息列表，单次同步最多1000条
	_styleBasicInfoBoList []StyleBasicInfoBo
}

// New
