package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabatclsfulfillqaordercreate 创单接口
// alibaba.tcls.fulfill.qa.order.create
//
// 根据历史测试履约单号，复制一个同样镜像的履约单号并下发给大润发仓（api实现已经限制了测试数据）
func Alibabatclsfulfillqaordercreate(clt *core.SDKClient, req *logistic.AlibabatclsfulfillqaordercreateAPIRequest, session string) (*logistic.AlibabatclsfulfillqaordercreateAPIResponse, error) {
	var resp logistic.AlibabatclsfulfillqaordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
