package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoWeikePerformancePut 提交客服绩效接口
// taobao.weike.performance.put
//
// 提交客服绩效接口
func TaobaoWeikePerformancePut(clt *core.SDKClient, req *servicecenter.TaobaoWeikePerformancePutAPIRequest, session string) (*servicecenter.TaobaoWeikePerformancePutAPIResponse, error) {
	var resp servicecenter.TaobaoWeikePerformancePutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
