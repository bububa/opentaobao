package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallMallitemcenterSupplierPriceUploadAPIResponse
天猫服务商服务报价上传 API返回值
tmall.mallitemcenter.supplier.price.upload

天猫服务商上传服务价格 */
type TmallMallitemcenterSupplierPriceUploadAPIResponse struct {
	model.CommonResponse
	TmallMallitemcenterSupplierPriceUploadAPIResponseModel
}

// TmallMallitemcenterSupplierPriceUploadAPIResponseModel is 天猫服务商服务报价上传 成功返回结果
type TmallMallitemcenterSupplierPriceUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mallitemcenter_supplier_price_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallMallitemcenterSupplierPriceUploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
