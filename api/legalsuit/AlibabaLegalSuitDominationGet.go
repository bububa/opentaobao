package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitdominationget 查询管辖信息
// alibaba.legal.suit.domination.get
//
// 查询管辖信息
func Alibabalegalsuitdominationget(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitdominationgetAPIRequest, session string) (*legalsuit.AlibabalegalsuitdominationgetAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitdominationgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
