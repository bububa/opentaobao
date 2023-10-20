package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangdeliverycreateupdate 新建/更新配资源
// alibaba.dchain.aoxiang.delivery.create.update
//
// 新建/更新配资源
func Alibabadchainaoxiangdeliverycreateupdate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangdeliverycreateupdateAPIRequest, session string) (*ascp.AlibabadchainaoxiangdeliverycreateupdateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangdeliverycreateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
