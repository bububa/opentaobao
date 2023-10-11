package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressAddressBlacklistTmsDelete 上门取退可揽范围黑名单删除接口
// taobao.logistics.express.address.blacklist.tms.delete
//
// 上门取退可揽范围黑名单删除接口
func TaobaoLogisticsExpressAddressBlacklistTmsDelete(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest, session string) (*ascp.TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse, error) {
	var resp ascp.TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
