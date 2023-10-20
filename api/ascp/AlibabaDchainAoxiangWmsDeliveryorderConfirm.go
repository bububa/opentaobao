package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangwmsdeliveryorderconfirm 回传发货单确认
// alibaba.dchain.aoxiang.wms.deliveryorder.confirm
//
// 回传发货单确认
func Alibabadchainaoxiangwmsdeliveryorderconfirm(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest, session string) (*ascp.AlibabadchainaoxiangwmsdeliveryorderconfirmAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangwmsdeliveryorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
