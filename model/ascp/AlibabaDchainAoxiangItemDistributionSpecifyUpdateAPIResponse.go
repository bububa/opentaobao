package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemdistributionspecifyupdateAPIResponse 指定分销商进行铺货(专享) - 修改 API返回值
// alibaba.dchain.aoxiang.item.distribution.specify.update
//
// 选择店铺的商品进行指定分销商铺货。 可以指定对应的分销商对应的价格
type AlibabadchainaoxiangitemdistributionspecifyupdateAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangitemdistributionspecifyupdateAPIResponseModel
}

// AlibabadchainaoxiangitemdistributionspecifyupdateAPIResponseModel is 指定分销商进行铺货(专享) - 修改 成功返回结果
type AlibabadchainaoxiangitemdistributionspecifyupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_item_distribution_specify_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	CreateItemDistributionResponse *TopResponse `json:"create_item_distribution_response,omitempty" xml:"create_item_distribution_response,omitempty"`
}
