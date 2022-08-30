package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemmappingDeleteAPIResponse 删除商货品关联关系 API返回值
// alibaba.dchain.aoxiang.itemmapping.delete
//
// 删除商货品关联关系
type AlibabaDchainAoxiangItemmappingDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangItemmappingDeleteAPIResponseModel
}

// AlibabaDchainAoxiangItemmappingDeleteAPIResponseModel is 删除商货品关联关系 成功返回结果
type AlibabaDchainAoxiangItemmappingDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_itemmapping_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	ItemMappingDeleteResponse *ItemMappingDeleteResponse `json:"item_mapping_delete_response,omitempty" xml:"item_mapping_delete_response,omitempty"`
}
