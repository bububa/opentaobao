package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoRdcAligeniusWarehouseReverseUploading 销退单上传
// taobao.rdc.aligenius.warehouse.reverse.uploading
//
// 主要用于商家上传仓库销退单信息
func TaobaoRdcAligeniusWarehouseReverseUploading(clt *core.SDKClient, req *logistic.TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest, session string) (*logistic.TaobaoRdcAligeniusWarehouseReverseUploadingAPIResponse, error) {
	var resp logistic.TaobaoRdcAligeniusWarehouseReverseUploadingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
