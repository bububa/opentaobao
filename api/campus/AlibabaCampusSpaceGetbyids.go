package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

/* AlibabaCampusSpaceGetbyids
根据ids和类型查询空间列表
alibaba.campus.space.getbyids

根据ids和类型查询空间列表 */
func AlibabaCampusSpaceGetbyids(clt *core.SDKClient, req *campus.AlibabaCampusSpaceGetbyidsAPIRequest, session string) (*campus.AlibabaCampusSpaceGetbyidsAPIResponse, error) {
	var resp campus.AlibabaCampusSpaceGetbyidsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
