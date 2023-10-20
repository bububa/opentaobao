package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Alibabaretailcommissionordersync 分佣数据传输
// alibaba.retail.commission.order.sync
//
// 同步分佣结果
func Alibabaretailcommissionordersync(clt *core.SDKClient, req *omniorder.AlibabaretailcommissionordersyncAPIRequest, session string) (*omniorder.AlibabaretailcommissionordersyncAPIResponse, error) {
	var resp omniorder.AlibabaretailcommissionordersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
