package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosstoregetstorelist 根据屏编号获取专柜集
// alibaba.mos.store.getstorelist
//
// 根据屏编号获取专柜集
func Alibabamosstoregetstorelist(clt *core.SDKClient, req *mos.AlibabamosstoregetstorelistAPIRequest, session string) (*mos.AlibabamosstoregetstorelistAPIResponse, error) {
	var resp mos.AlibabamosstoregetstorelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
