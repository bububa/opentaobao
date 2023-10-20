package fivee

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeCompanyGetAPIResponse 查询商信息 API返回值
// taobao.fivee.company.get
//
// 资质共享平台查询商信息
type TaobaoFiveeCompanyGetAPIResponse struct {
	model.CommonResponse
	TaobaoFiveeCompanyGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFiveeCompanyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFiveeCompanyGetAPIResponseModel).Reset()
}

// TaobaoFiveeCompanyGetAPIResponseModel is 查询商信息 成功返回结果
type TaobaoFiveeCompanyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fivee_company_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFiveeCompanyGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFiveeCompanyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFiveeCompanyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFiveeCompanyGetAPIResponse)
	},
}

// GetTaobaoFiveeCompanyGetAPIResponse 从 sync.Pool 获取 TaobaoFiveeCompanyGetAPIResponse
func GetTaobaoFiveeCompanyGetAPIResponse() *TaobaoFiveeCompanyGetAPIResponse {
	return poolTaobaoFiveeCompanyGetAPIResponse.Get().(*TaobaoFiveeCompanyGetAPIResponse)
}

// ReleaseTaobaoFiveeCompanyGetAPIResponse 将 TaobaoFiveeCompanyGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFiveeCompanyGetAPIResponse(v *TaobaoFiveeCompanyGetAPIResponse) {
	v.Reset()
	poolTaobaoFiveeCompanyGetAPIResponse.Put(v)
}
