package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkseriessort 系列品-商品排序
// alibaba.wdk.series.sort
//
// 系列品商品变更-商品排序
func Alibabawdkseriessort(clt *core.SDKClient, req *wdk.AlibabawdkseriessortAPIRequest, session string) (*wdk.AlibabawdkseriessortAPIResponse, error) {
	var resp wdk.AlibabawdkseriessortAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
