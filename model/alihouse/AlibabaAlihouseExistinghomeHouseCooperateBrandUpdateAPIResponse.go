package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse 租房合作品牌更新接口 API返回值
// alibaba.alihouse.existinghome.house.cooperate.brand.update
//
// 租房合作品牌更新接口
type AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponseModel
}

// AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponseModel is 租房合作品牌更新接口 成功返回结果
type AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_cooperate_brand_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值对象
	Result *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
