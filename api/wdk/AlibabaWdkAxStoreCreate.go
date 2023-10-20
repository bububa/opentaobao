package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkaxstorecreate 翱象经营店创建接口
// alibaba.wdk.ax.store.create
//
// 翱象经营店创建
func Alibabawdkaxstorecreate(clt *core.SDKClient, req *wdk.AlibabawdkaxstorecreateAPIRequest, session string) (*wdk.AlibabawdkaxstorecreateAPIResponse, error) {
	var resp wdk.AlibabawdkaxstorecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
