package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// Alibabalegalcaseentrustcallback 委托回调接口
// alibaba.legal.case.entrust.callback
//
// 委托回调接口
func Alibabalegalcaseentrustcallback(clt *core.SDKClient, req *legalcase.AlibabalegalcaseentrustcallbackAPIRequest, session string) (*legalcase.AlibabalegalcaseentrustcallbackAPIResponse, error) {
	var resp legalcase.AlibabalegalcaseentrustcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
