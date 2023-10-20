package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMallitemcenterSupplierPriceUpload 天猫服务商服务报价上传
// tmall.mallitemcenter.supplier.price.upload
//
// 天猫服务商上传服务价格
func TmallMallitemcenterSupplierPriceUpload(clt *core.SDKClient, req *tmallservice.TmallMallitemcenterSupplierPriceUploadAPIRequest, session string) (*tmallservice.TmallMallitemcenterSupplierPriceUploadAPIResponse, error) {
	var resp tmallservice.TmallMallitemcenterSupplierPriceUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
