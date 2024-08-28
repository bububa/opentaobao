package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoRdcAligeniusWarehouseReverseUploading 销退单上传
// taobao.rdc.aligenius.warehouse.reverse.uploading
//
// 主要用于商家上传仓库销退单信息
func TaobaoRdcAligeniusWarehouseReverseUploading(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest, resp *logistic.TaobaoRdcAligeniusWarehouseReverseUploadingAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
