package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangwmsdeliveryordercreate 回传仓库接单通知
// alibaba.dchain.aoxiang.wms.deliveryorder.create
//
// WMS上报仓库接单节点状态信息，代表接单环节。
func Alibabadchainaoxiangwmsdeliveryordercreate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest, session string) (*ascp.AlibabadchainaoxiangwmsdeliveryordercreateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangwmsdeliveryordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
