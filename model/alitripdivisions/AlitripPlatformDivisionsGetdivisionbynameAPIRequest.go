package alitripdivisions

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripPlatformDivisionsGetdivisionbynameAPIRequest
根据中文名称与行政区划级别查询行政区划数据 API请求
alitrip.platform.divisions.getdivisionbyname

根据中文名称与行政区划级别查询行政区划数据 */
type AlitripPlatformDivisionsGetdivisionbynameAPIRequest struct {
	model.Params
	// 行政区划名称
	_name string
	// 行政区划级别ALL(0, "全部"),  	CONTINENT(1, "大洲"),  	COUNTRY(2, "国家"),  	PROVINCE(3, "省份"),  	CITY(4, "城市"),  	DISTRICT(5, "区县"),  	STREET(6, "街道")
	_level int64
}

// New
