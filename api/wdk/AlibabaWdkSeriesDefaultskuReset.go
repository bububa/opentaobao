package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkSeriesDefaultskuReset
系列品商品变更-重置默认商品
alibaba.wdk.series.defaultsku.reset

系列品商品变更-重置默认商品 */
func AlibabaWdkSeriesDefaultskuReset(clt *core.SDKClient, req *wdk.AlibabaWdkSeriesDefaultskuResetAPIRequest, session string) (*wdk.AlibabaWdkSeriesDefaultskuResetAPIResponse, error) {
	var resp wdk.AlibabaWdkSeriesDefaultskuResetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
