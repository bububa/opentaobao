package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthimriskquery 问诊质控接口
// alibaba.alihealth.imrisk.query
//
// 阿里健康的问诊质控接口
func Alibabaalihealthimriskquery(clt *core.SDKClient, req *alihealth2.AlibabaalihealthimriskqueryAPIRequest, session string) (*alihealth2.AlibabaalihealthimriskqueryAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthimriskqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
