package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapurcreatedo top创建DO/RT接口
// alibaba.pur.create.do
//
// 创建发货单,先创建DO，异步创建RT
func Alibabapurcreatedo(clt *core.SDKClient, req *pur.AlibabapurcreatedoAPIRequest, session string) (*pur.AlibabapurcreatedoAPIResponse, error) {
	var resp pur.AlibabapurcreatedoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
