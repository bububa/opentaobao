package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSeriesCreate 系列品变更-新增系列
// alibaba.wdk.series.create
//
// 系列品变更-新增系列
func AlibabaWdkSeriesCreate(clt *core.SDKClient, req *wdk.AlibabaWdkSeriesCreateAPIRequest, session string) (*wdk.AlibabaWdkSeriesCreateAPIResponse, error) {
	var resp wdk.AlibabaWdkSeriesCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
