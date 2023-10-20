package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaalicomvtdistributequeryprotocol 通信业务外放协议查询
// alibaba.alicom.vt.distribute.queryprotocol
//
// 通信业务外放协议查询
func Alibabaalicomvtdistributequeryprotocol(clt *core.SDKClient, req *alicom.AlibabaalicomvtdistributequeryprotocolAPIRequest, session string) (*alicom.AlibabaalicomvtdistributequeryprotocolAPIResponse, error) {
	var resp alicom.AlibabaalicomvtdistributequeryprotocolAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
