package alitripdivisions

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripdivisions"
)

/* AlitripPlatformDivisionsQuerybyparentid
根据父节点id查询下级行政区划数据
alitrip.platform.divisions.querybyparentid

根据行政区划id查询下一层级行政区划数据 */
func AlitripPlatformDivisionsQuerybyparentid(clt *core.SDKClient, req *alitripdivisions.AlitripPlatformDivisionsQuerybyparentidAPIRequest, session string) (*alitripdivisions.AlitripPlatformDivisionsQuerybyparentidAPIResponse, error) {
	var resp alitripdivisions.AlitripPlatformDivisionsQuerybyparentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
