package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemmappingupdateasyncAPIResponse 创建/更新商货品关联关系 API返回值
// alibaba.dchain.aoxiang.itemmapping.update.async
//
// 创建/更新商货品关联关系
type AlibabadchainaoxiangitemmappingupdateasyncAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangitemmappingupdateasyncAPIResponseModel
}

// AlibabadchainaoxiangitemmappingupdateasyncAPIResponseModel is 创建/更新商货品关联关系 成功返回结果
type AlibabadchainaoxiangitemmappingupdateasyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_itemmapping_update_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	ItemMappingUpdateResponse *ItemMappingUpdateAsyncResponse `json:"item_mapping_update_response,omitempty" xml:"item_mapping_update_response,omitempty"`
}
