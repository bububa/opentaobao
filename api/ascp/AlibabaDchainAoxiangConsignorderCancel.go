package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangConsignorderCancel 自动流转单据取消仓内发货
// alibaba.dchain.aoxiang.consignorder.cancel
//
// 自动流转单据取消仓内发货
func AlibabaDchainAoxiangConsignorderCancel(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangConsignorderCancelAPIRequest, session string) (*ascp.AlibabaDchainAoxiangConsignorderCancelAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangConsignorderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
