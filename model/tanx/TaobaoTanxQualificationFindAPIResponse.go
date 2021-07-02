package tanx

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxQualificationFindAPIResponse 资质查询接口 API返回值
// taobao.tanx.qualification.find
//
// 资质查询接口
type TaobaoTanxQualificationFindAPIResponse struct {
	model.CommonResponse
	TaobaoTanxQualificationFindAPIResponseModel
}

// TaobaoTanxQualificationFindAPIResponseModel is 资质查询接口 成功返回结果
type TaobaoTanxQualificationFindAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_qualification_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 返回的资质内容dto
	QualificationList []QualificationDto `json:"qualification_list,omitempty" xml:"qualification_list>qualification_dto,omitempty"`
	// 查询返回总条数
	Count string `json:"count,omitempty" xml:"count,omitempty"`
}
