package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkumsoutboundsortingscancontainer dps分货-扫描分货容器判断是否可用
// wdk.ums.outbound.sorting.scancontainer
//
// dps分货-扫描分货容器判断是否可用
func Wdkumsoutboundsortingscancontainer(clt *core.SDKClient, req *wdk.WdkumsoutboundsortingscancontainerAPIRequest, session string) (*wdk.WdkumsoutboundsortingscancontainerAPIResponse, error) {
	var resp wdk.WdkumsoutboundsortingscancontainerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
