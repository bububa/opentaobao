package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaAlihealthNrRxQueryimage o2o查看处方图片
// alibaba.alihealth.nr.rx.queryimage
//
// o2o商家查看处方图片，包括电子图片与纸质图片
func AlibabaAlihealthNrRxQueryimage(clt *core.SDKClient, req *drug.AlibabaAlihealthNrRxQueryimageAPIRequest, resp *drug.AlibabaAlihealthNrRxQueryimageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
