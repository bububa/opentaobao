package tanx

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxQualificationAdvertiserAddAPIResponse 新增广告主接口 API返回值
// taobao.tanx.qualification.advertiser.add
//
// 外部dsp调用接口时会根据广告主名称和类型在tanx系统中新增一个广告主
type TaobaoTanxQualificationAdvertiserAddAPIResponse struct {
	model.CommonResponse
	TaobaoTanxQualificationAdvertiserAddAPIResponseModel
}

// TaobaoTanxQualificationAdvertiserAddAPIResponseModel is 新增广告主接口 成功返回结果
type TaobaoTanxQualificationAdvertiserAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_qualification_advertiser_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 返回的广告主dto对象
	AdvertiserList []AdvertiserDto `json:"advertiser_list,omitempty" xml:"advertiser_list>advertiser_dto,omitempty"`
}
