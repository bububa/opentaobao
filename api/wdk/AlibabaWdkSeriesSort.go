package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSeriesSort 系列品-商品排序
// alibaba.wdk.series.sort
//
// 系列品商品变更-商品排序
func AlibabaWdkSeriesSort(clt *core.SDKClient, req *wdk.AlibabaWdkSeriesSortAPIRequest, session string) (*wdk.AlibabaWdkSeriesSortAPIResponse, error) {
	var resp wdk.AlibabaWdkSeriesSortAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
