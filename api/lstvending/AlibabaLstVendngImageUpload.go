package lstvending

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

/* AlibabaLstVendngImageUpload
售货机商品图片上传
alibaba.lst.vendng.image.upload

零售通自动售货机商品图片上传接口，主要为ISV厂商提供图片同步的通道，从而建立统一的商品图片库。 */
func AlibabaLstVendngImageUpload(clt *core.SDKClient, req *lstvending.AlibabaLstVendngImageUploadAPIRequest, session string) (*lstvending.AlibabaLstVendngImageUploadAPIResponse, error) {
	var resp lstvending.AlibabaLstVendngImageUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
