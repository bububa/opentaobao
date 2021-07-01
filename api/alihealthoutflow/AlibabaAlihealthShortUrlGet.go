package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* AlibabaAlihealthShortUrlGet
支付宝短链跳转三方h5通用接口
alibaba.alihealth.short.url.get

支付宝短链跳转三方h5通用接口 */
func AlibabaAlihealthShortUrlGet(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthShortUrlGetAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthShortUrlGetAPIResponse, error) {
	var resp alihealthoutflow.AlibabaAlihealthShortUrlGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
