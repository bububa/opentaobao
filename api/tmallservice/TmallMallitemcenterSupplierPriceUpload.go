package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallmallitemcentersupplierpriceupload 天猫服务商服务报价上传
// tmall.mallitemcenter.supplier.price.upload
//
// 天猫服务商上传服务价格
func Tmallmallitemcentersupplierpriceupload(clt *core.SDKClient, req *tmallservice.TmallmallitemcentersupplierpriceuploadAPIRequest, session string) (*tmallservice.TmallmallitemcentersupplierpriceuploadAPIResponse, error) {
	var resp tmallservice.TmallmallitemcentersupplierpriceuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
