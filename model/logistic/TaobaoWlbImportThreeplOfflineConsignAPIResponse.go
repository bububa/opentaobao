package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbimportthreeplofflineconsignAPIResponse 3PL直邮线下发货 API返回值
// taobao.wlb.import.threepl.offline.consign
//
// 菜鸟认证直邮线下发货
type TaobaowlbimportthreeplofflineconsignAPIResponse struct {
	model.CommonResponse
	TaobaowlbimportthreeplofflineconsignAPIResponseModel
}

// TaobaowlbimportthreeplofflineconsignAPIResponseModel is 3PL直邮线下发货 成功返回结果
type TaobaowlbimportthreeplofflineconsignAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_import_threepl_offline_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaowlbimportthreeplofflineconsignTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
