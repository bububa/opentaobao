package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkbmpaiyangstatdataquery 派样统计数据查询
// alibaba.wdk.bm.paiyang.stat.data.query
//
// 派样统计数据查询
func Alibabawdkbmpaiyangstatdataquery(clt *core.SDKClient, req *wdk.AlibabawdkbmpaiyangstatdataqueryAPIRequest, session string) (*wdk.AlibabawdkbmpaiyangstatdataqueryAPIResponse, error) {
	var resp wdk.AlibabawdkbmpaiyangstatdataqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
