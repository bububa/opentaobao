package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// Alibabalegalcaseentrustget 委托
// alibaba.legal.case.entrust.get
//
// 获取委托案件的基本信息
func Alibabalegalcaseentrustget(clt *core.SDKClient, req *legalcase.AlibabalegalcaseentrustgetAPIRequest, session string) (*legalcase.AlibabalegalcaseentrustgetAPIResponse, error) {
	var resp legalcase.AlibabalegalcaseentrustgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
