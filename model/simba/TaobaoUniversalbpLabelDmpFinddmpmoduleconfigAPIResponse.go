package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse 查询dmp浮层配置 API返回值
// taobao.universalbp.label.dmp.finddmpmoduleconfig
//
// 入参账号信息，出参达摩盘相关配置信息
type TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponseModel).Reset()
}

// TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponseModel is 查询dmp浮层配置 成功返回结果
type TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_label_dmp_finddmpmoduleconfig_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse)
	},
}

// GetTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse
func GetTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse() *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse {
	return poolTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse.Get().(*TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse)
}

// ReleaseTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse 将 TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse(v *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse.Put(v)
}
