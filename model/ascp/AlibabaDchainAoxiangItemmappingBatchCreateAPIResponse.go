package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemmappingbatchcreateAPIResponse 新建商货品关联 API返回值
// alibaba.dchain.aoxiang.itemmapping.batch.create
//
// 新建商货品关联
type AlibabadchainaoxiangitemmappingbatchcreateAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangitemmappingbatchcreateAPIResponseModel
}

// AlibabadchainaoxiangitemmappingbatchcreateAPIResponseModel is 新建商货品关联 成功返回结果
type AlibabadchainaoxiangitemmappingbatchcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_itemmapping_batch_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchCreateItemMappingResponse *BatchCreateItemMappingResponse `json:"batch_create_item_mapping_response,omitempty" xml:"batch_create_item_mapping_response,omitempty"`
}
