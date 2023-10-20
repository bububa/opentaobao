package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemdistributioncreateAPIResponse 选择店铺商品并进行铺货（通用） API返回值
// alibaba.dchain.aoxiang.item.distribution.create
//
// 选择店铺商品并进行铺货, 铺货给所有的合作分销商。设定的价格为通用价格
type AlibabadchainaoxiangitemdistributioncreateAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangitemdistributioncreateAPIResponseModel
}

// AlibabadchainaoxiangitemdistributioncreateAPIResponseModel is 选择店铺商品并进行铺货（通用） 成功返回结果
type AlibabadchainaoxiangitemdistributioncreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_item_distribution_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	CreateItemDistributionResponse *TopResponse `json:"create_item_distribution_response,omitempty" xml:"create_item_distribution_response,omitempty"`
}
