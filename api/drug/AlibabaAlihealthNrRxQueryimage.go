package drug

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaAlihealthNrRxQueryimage o2o查看处方图片
// alibaba.alihealth.nr.rx.queryimage
//
// o2o商家查看处方图片，包括电子图片与纸质图片
func AlibabaAlihealthNrRxQueryimage(ctx context.Context, clt *core.SDKClient, req *drug.AlibabaAlihealthNrRxQueryimageAPIRequest, resp *drug.AlibabaAlihealthNrRxQueryimageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
