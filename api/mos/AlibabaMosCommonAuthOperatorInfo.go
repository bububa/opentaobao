package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamoscommonauthoperatorinfo 获取当前人员信息
// alibaba.mos.common.auth.operator.info
//
// 获取当前人员信息
func Alibabamoscommonauthoperatorinfo(clt *core.SDKClient, req *mos.AlibabamoscommonauthoperatorinfoAPIRequest, session string) (*mos.AlibabamoscommonauthoperatorinfoAPIResponse, error) {
	var resp mos.AlibabamoscommonauthoperatorinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
