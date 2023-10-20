package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// Alibabaalihealthnrrxqueryimage o2o查看处方图片
// alibaba.alihealth.nr.rx.queryimage
//
// o2o商家查看处方图片，包括电子图片与纸质图片
func Alibabaalihealthnrrxqueryimage(clt *core.SDKClient, req *drug.AlibabaalihealthnrrxqueryimageAPIRequest, session string) (*drug.AlibabaalihealthnrrxqueryimageAPIResponse, error) {
	var resp drug.AlibabaalihealthnrrxqueryimageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
