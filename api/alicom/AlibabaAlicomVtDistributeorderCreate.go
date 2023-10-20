package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaalicomvtdistributeordercreate 通信业务外放下单
// alibaba.alicom.vt.distributeorder.create
//
// 通信业务外放下单接口
func Alibabaalicomvtdistributeordercreate(clt *core.SDKClient, req *alicom.AlibabaalicomvtdistributeordercreateAPIRequest, session string) (*alicom.AlibabaalicomvtdistributeordercreateAPIResponse, error) {
	var resp alicom.AlibabaalicomvtdistributeordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
