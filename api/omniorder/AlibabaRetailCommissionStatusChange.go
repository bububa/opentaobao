package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Alibabaretailcommissionstatuschange 分佣状态变更
// alibaba.retail.commission.status.change
//
// 分佣系统，分佣状态变更接口
func Alibabaretailcommissionstatuschange(clt *core.SDKClient, req *omniorder.AlibabaretailcommissionstatuschangeAPIRequest, session string) (*omniorder.AlibabaretailcommissionstatuschangeAPIResponse, error) {
	var resp omniorder.AlibabaretailcommissionstatuschangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
