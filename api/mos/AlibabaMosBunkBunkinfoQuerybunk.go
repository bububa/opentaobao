package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosbunkbunkinfoquerybunk 根据合同号查询铺位信息
// alibaba.mos.bunk.bunkinfo.querybunk
//
// 根据合同号查询铺位信息
func Alibabamosbunkbunkinfoquerybunk(clt *core.SDKClient, req *mos.AlibabamosbunkbunkinfoquerybunkAPIRequest, session string) (*mos.AlibabamosbunkbunkinfoquerybunkAPIResponse, error) {
	var resp mos.AlibabamosbunkbunkinfoquerybunkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
