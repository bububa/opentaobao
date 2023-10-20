package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangphysicsinventorybatchuploadasync 实仓库存同步
// alibaba.dchain.aoxiang.physics.inventory.batch.upload.async
//
// 实仓库存同步
func Alibabadchainaoxiangphysicsinventorybatchuploadasync(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest, session string) (*ascp.AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
