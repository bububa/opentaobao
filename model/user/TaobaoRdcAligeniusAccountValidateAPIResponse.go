package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusAccountValidateAPIResponse AG商家账号校验 API返回值
// taobao.rdc.aligenius.account.validate
//
// 提供应对接AG的erp系统查询其旗下的商家是否为AG商家
type TaobaoRdcAligeniusAccountValidateAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusAccountValidateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusAccountValidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRdcAligeniusAccountValidateAPIResponseModel).Reset()
}

// TaobaoRdcAligeniusAccountValidateAPIResponseModel is AG商家账号校验 成功返回结果
type TaobaoRdcAligeniusAccountValidateAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_account_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoRdcAligeniusAccountValidateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusAccountValidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRdcAligeniusAccountValidateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusAccountValidateAPIResponse)
	},
}

// GetTaobaoRdcAligeniusAccountValidateAPIResponse 从 sync.Pool 获取 TaobaoRdcAligeniusAccountValidateAPIResponse
func GetTaobaoRdcAligeniusAccountValidateAPIResponse() *TaobaoRdcAligeniusAccountValidateAPIResponse {
	return poolTaobaoRdcAligeniusAccountValidateAPIResponse.Get().(*TaobaoRdcAligeniusAccountValidateAPIResponse)
}

// ReleaseTaobaoRdcAligeniusAccountValidateAPIResponse 将 TaobaoRdcAligeniusAccountValidateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRdcAligeniusAccountValidateAPIResponse(v *TaobaoRdcAligeniusAccountValidateAPIResponse) {
	v.Reset()
	poolTaobaoRdcAligeniusAccountValidateAPIResponse.Put(v)
}
