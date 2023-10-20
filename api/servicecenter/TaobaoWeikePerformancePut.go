package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaoweikeperformanceput 提交客服绩效接口
// taobao.weike.performance.put
//
// 提交客服绩效接口
func Taobaoweikeperformanceput(clt *core.SDKClient, req *servicecenter.TaobaoweikeperformanceputAPIRequest, session string) (*servicecenter.TaobaoweikeperformanceputAPIResponse, error) {
	var resp servicecenter.TaobaoweikeperformanceputAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
