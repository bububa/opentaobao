package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitcaseget 获取案件信息接口v2版本
// alibaba.legal.suit.case.get
//
// 获取案件信息
func Alibabalegalsuitcaseget(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitcasegetAPIRequest, session string) (*legalsuit.AlibabalegalsuitcasegetAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitcasegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
