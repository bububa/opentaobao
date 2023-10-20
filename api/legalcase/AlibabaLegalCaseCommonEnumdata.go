package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// Alibabalegalcasecommonenumdata 获取通用枚举值接口
// alibaba.legal.case.common.enumdata
//
// 获取通用枚举值接口
func Alibabalegalcasecommonenumdata(clt *core.SDKClient, req *legalcase.AlibabalegalcasecommonenumdataAPIRequest, session string) (*legalcase.AlibabalegalcasecommonenumdataAPIResponse, error) {
	var resp legalcase.AlibabalegalcasecommonenumdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
