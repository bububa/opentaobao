package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// Alibabalegalcasecommonnotice 消息通知
// alibaba.legal.case.common.notice
//
// 同步通知给诉讼系统
func Alibabalegalcasecommonnotice(clt *core.SDKClient, req *legalcase.AlibabalegalcasecommonnoticeAPIRequest, session string) (*legalcase.AlibabalegalcasecommonnoticeAPIResponse, error) {
	var resp legalcase.AlibabalegalcasecommonnoticeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
