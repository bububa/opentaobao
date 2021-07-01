package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFiveeImportproductPublishAPIRequest
进口商品发布 API请求
taobao.fivee.importproduct.publish

直营业务商家入住发布商品时，上传商品及商家证照信息 */
type TaobaoFiveeImportproductPublishAPIRequest struct {
	model.Params
	// 进口商品
	_importProduct *ImportProduct
	// bu身份标识
	_paramBucode string
}

// New
