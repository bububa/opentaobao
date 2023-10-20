package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkaxstoreupdate 翱翔经营店更新接口
// alibaba.wdk.ax.store.update
//
// 翱翔经营店更新接口
func Alibabawdkaxstoreupdate(clt *core.SDKClient, req *wdk.AlibabawdkaxstoreupdateAPIRequest, session string) (*wdk.AlibabawdkaxstoreupdateAPIResponse, error) {
	var resp wdk.AlibabawdkaxstoreupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
