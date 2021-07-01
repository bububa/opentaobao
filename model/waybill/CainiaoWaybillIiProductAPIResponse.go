package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiProductAPIResponse
商家查询物流商产品类型接口 API返回值
cainiao.waybill.ii.product

商家可以查询物流商的产品类型和服务能力。 */
type CainiaoWaybillIiProductAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiProductAPIResponseModel
}

// CainiaoWaybillIiProductAPIResponseModel is 商家查询物流商产品类型接口 成功返回结果
type CainiaoWaybillIiProductAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	ProductTypes []WaybillProductType `json:"product_types,omitempty" xml:"product_types>waybill_product_type,omitempty"`
}
