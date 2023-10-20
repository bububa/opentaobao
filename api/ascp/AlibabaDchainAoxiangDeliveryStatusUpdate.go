package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangdeliverystatusupdate 启用/停用配资源
// alibaba.dchain.aoxiang.delivery.status.update
//
// 启用/停用配资源
func Alibabadchainaoxiangdeliverystatusupdate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangdeliverystatusupdateAPIRequest, session string) (*ascp.AlibabadchainaoxiangdeliverystatusupdateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangdeliverystatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
