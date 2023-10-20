package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytquerybatchprodAPIResponse 批次产品查询(根据企业名和批次号查询产品信息) API返回值
// alibaba.alihealth.drug.kyt.querybatchprod
//
// 根据企业名和批次号查询药品信息，支持使用更名之前的老企业名查询，支持批次号大小写模糊，应用于药店或医院入库环节，通过在入库环节获取赋码的产品目录，可强制要求对相应的产品必须进行扫码入库；
type AlibabaalihealthdrugkytquerybatchprodAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytquerybatchprodAPIResponseModel
}

// AlibabaalihealthdrugkytquerybatchprodAPIResponseModel is 批次产品查询(根据企业名和批次号查询产品信息) 成功返回结果
type AlibabaalihealthdrugkytquerybatchprodAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_querybatchprod_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *AlibabaalihealthdrugkytquerybatchprodResult `json:"result,omitempty" xml:"result,omitempty"`
}
