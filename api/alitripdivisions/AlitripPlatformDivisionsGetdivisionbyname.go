package alitripdivisions

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripdivisions"
)

// Alitripplatformdivisionsgetdivisionbyname 根据中文名称与行政区划级别查询行政区划数据
// alitrip.platform.divisions.getdivisionbyname
//
// 根据中文名称与行政区划级别查询行政区划数据
func Alitripplatformdivisionsgetdivisionbyname(clt *core.SDKClient, req *alitripdivisions.AlitripplatformdivisionsgetdivisionbynameAPIRequest, session string) (*alitripdivisions.AlitripplatformdivisionsgetdivisionbynameAPIResponse, error) {
	var resp alitripdivisions.AlitripplatformdivisionsgetdivisionbynameAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
