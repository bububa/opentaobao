package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse 后端商品库存占用调整接口 API返回值
// taobao.jst.astrolabe.storeinventory.adjust
//
// 当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
type TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse struct {
	model.CommonResponse
	TaobaoJstAstrolabeStoreinventoryAdjustAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstAstrolabeStoreinventoryAdjustAPIResponseModel).Reset()
}

// TaobaoJstAstrolabeStoreinventoryAdjustAPIResponseModel is 后端商品库存占用调整接口 成功返回结果
type TaobaoJstAstrolabeStoreinventoryAdjustAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_storeinventory_adjust_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应标签
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应编码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryAdjustAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Flag = ""
	m.ResultCode = ""
	m.Message = ""
}

var poolTaobaoJstAstrolabeStoreinventoryAdjustAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse)
	},
}

// GetTaobaoJstAstrolabeStoreinventoryAdjustAPIResponse 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse
func GetTaobaoJstAstrolabeStoreinventoryAdjustAPIResponse() *TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse {
	return poolTaobaoJstAstrolabeStoreinventoryAdjustAPIResponse.Get().(*TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryAdjustAPIResponse 将 TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryAdjustAPIResponse(v *TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryAdjustAPIResponse.Put(v)
}
