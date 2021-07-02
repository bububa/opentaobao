package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkStockPublish 五道口库存发布接口（针对外部渠道）
// alibaba.wdk.stock.publish
//
// 五道口库存发布接口（针对外部渠道）
func AlibabaWdkStockPublish(clt *core.SDKClient, req *wdk.AlibabaWdkStockPublishAPIRequest, session string) (*wdk.AlibabaWdkStockPublishAPIResponse, error) {
	var resp wdk.AlibabaWdkStockPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
