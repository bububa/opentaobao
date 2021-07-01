package alitripdivisions

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripPlatformDivisionsQuerybyparentidAPIRequest
根据父节点id查询下级行政区划数据 API请求
alitrip.platform.divisions.querybyparentid

根据行政区划id查询下一层级行政区划数据 */
type AlitripPlatformDivisionsQuerybyparentidAPIRequest struct {
	model.Params
	// 行政区划父id
	_paramLong int64
}

// New
