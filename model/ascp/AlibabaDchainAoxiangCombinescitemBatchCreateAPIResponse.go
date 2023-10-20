package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangcombinescitembatchcreateAPIResponse 新建组合货品 API返回值
// alibaba.dchain.aoxiang.combinescitem.batch.create
//
// 新建组合货品
type AlibabadchainaoxiangcombinescitembatchcreateAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangcombinescitembatchcreateAPIResponseModel
}

// AlibabadchainaoxiangcombinescitembatchcreateAPIResponseModel is 新建组合货品 成功返回结果
type AlibabadchainaoxiangcombinescitembatchcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_combinescitem_batch_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchCreateCombineScitemResponse *BatchCreateCombineScItemResponse `json:"batch_create_combine_scitem_response,omitempty" xml:"batch_create_combine_scitem_response,omitempty"`
}
