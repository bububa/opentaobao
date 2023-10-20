package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsdeliverylinebatchdelete 线路能力删除
// taobao.logistics.delivery.line.batch.delete
//
// 线路能力删除
func Taobaologisticsdeliverylinebatchdelete(clt *core.SDKClient, req *ascp.TaobaologisticsdeliverylinebatchdeleteAPIRequest, session string) (*ascp.TaobaologisticsdeliverylinebatchdeleteAPIResponse, error) {
	var resp ascp.TaobaologisticsdeliverylinebatchdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
