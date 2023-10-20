package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaordcaligeniuswarehousereverseuploading 销退单上传
// taobao.rdc.aligenius.warehouse.reverse.uploading
//
// 主要用于商家上传仓库销退单信息
func Taobaordcaligeniuswarehousereverseuploading(clt *core.SDKClient, req *logistic.TaobaordcaligeniuswarehousereverseuploadingAPIRequest, session string) (*logistic.TaobaordcaligeniuswarehousereverseuploadingAPIResponse, error) {
	var resp logistic.TaobaordcaligeniuswarehousereverseuploadingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
