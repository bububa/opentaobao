package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

/* AlibabaPurPrCreate
下pr单
alibaba.pur.pr.create

下pr单 */
func AlibabaPurPrCreate(clt *core.SDKClient, req *pur.AlibabaPurPrCreateAPIRequest, session string) (*pur.AlibabaPurPrCreateAPIResponse, error) {
	var resp pur.AlibabaPurPrCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
