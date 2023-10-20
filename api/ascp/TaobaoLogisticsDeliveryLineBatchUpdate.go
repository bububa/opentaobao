package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsdeliverylinebatchupdate 线路能力创建/更新
// taobao.logistics.delivery.line.batch.update
//
// 线路能力创建/更新
func Taobaologisticsdeliverylinebatchupdate(clt *core.SDKClient, req *ascp.TaobaologisticsdeliverylinebatchupdateAPIRequest, session string) (*ascp.TaobaologisticsdeliverylinebatchupdateAPIResponse, error) {
	var resp ascp.TaobaologisticsdeliverylinebatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
