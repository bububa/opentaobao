package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitcourtentrustget 委托开庭服务查询
// alibaba.legal.suit.court.entrust.get
//
// 查询委托开庭信息
func Alibabalegalsuitcourtentrustget(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitcourtentrustgetAPIRequest, session string) (*legalsuit.AlibabalegalsuitcourtentrustgetAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitcourtentrustgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
