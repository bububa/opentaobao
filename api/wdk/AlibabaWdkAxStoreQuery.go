package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkaxstorequery 翱象经营店查询接口
// alibaba.wdk.ax.store.query
//
// 翱象经营店查询接口
func Alibabawdkaxstorequery(clt *core.SDKClient, req *wdk.AlibabawdkaxstorequeryAPIRequest, session string) (*wdk.AlibabawdkaxstorequeryAPIResponse, error) {
	var resp wdk.AlibabawdkaxstorequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
