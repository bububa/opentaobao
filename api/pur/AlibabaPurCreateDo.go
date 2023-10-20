package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurCreateDo top创建DO/RT接口
// alibaba.pur.create.do
//
// 创建发货单,先创建DO，异步创建RT
func AlibabaPurCreateDo(clt *core.SDKClient, req *pur.AlibabaPurCreateDoAPIRequest, session string) (*pur.AlibabaPurCreateDoAPIResponse, error) {
	var resp pur.AlibabaPurCreateDoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
