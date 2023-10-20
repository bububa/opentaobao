package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangscitembatchcreateAPIResponse 新建货品 API返回值
// alibaba.dchain.aoxiang.scitem.batch.create
//
// 新建货品
type AlibabadchainaoxiangscitembatchcreateAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangscitembatchcreateAPIResponseModel
}

// AlibabadchainaoxiangscitembatchcreateAPIResponseModel is 新建货品 成功返回结果
type AlibabadchainaoxiangscitembatchcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_scitem_batch_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchCreateScitemResponse *BatchCreateScItemResponse `json:"batch_create_scitem_response,omitempty" xml:"batch_create_scitem_response,omitempty"`
}
