package tanx

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoTanxQualificationFindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTanxQualificationFindAPIResponseModel).Reset()
}

// TaobaoTanxQualificationFindAPIResponseModel is 资质查询接口 成功返回结果
type TaobaoTanxQualificationFindAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_qualification_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的资质内容dto
	QualificationList []QualificationDto `json:"qualification_list,omitempty" xml:"qualification_list>qualification_dto,omitempty"`
	// 查询返回总条数
	Count string `json:"count,omitempty" xml:"count,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTanxQualificationFindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.QualificationList = m.QualificationList[:0]
	m.Count = ""
	m.IsSuccess = false
}

var poolTaobaoTanxQualificationFindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTanxQualificationFindAPIResponse)
	},
}

// GetTaobaoTanxQualificationFindAPIResponse 从 sync.Pool 获取 TaobaoTanxQualificationFindAPIResponse
func GetTaobaoTanxQualificationFindAPIResponse() *TaobaoTanxQualificationFindAPIResponse {
	return poolTaobaoTanxQualificationFindAPIResponse.Get().(*TaobaoTanxQualificationFindAPIResponse)
}

// ReleaseTaobaoTanxQualificationFindAPIResponse 将 TaobaoTanxQualificationFindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTanxQualificationFindAPIResponse(v *TaobaoTanxQualificationFindAPIResponse) {
	v.Reset()
	poolTaobaoTanxQualificationFindAPIResponse.Put(v)
}
