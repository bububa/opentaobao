package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpStdcategoryFindlistAPIResponse 人群相关类目查询 API返回值
// taobao.universalbp.stdcategory.findlist
//
// 入参商品信息，出参商品所属类别
type TaobaoUniversalbpStdcategoryFindlistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpStdcategoryFindlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpStdcategoryFindlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpStdcategoryFindlistAPIResponseModel).Reset()
}

// TaobaoUniversalbpStdcategoryFindlistAPIResponseModel is 人群相关类目查询 成功返回结果
type TaobaoUniversalbpStdcategoryFindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_stdcategory_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpStdcategoryFindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpStdcategoryFindlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpStdcategoryFindlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpStdcategoryFindlistAPIResponse)
	},
}

// GetTaobaoUniversalbpStdcategoryFindlistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpStdcategoryFindlistAPIResponse
func GetTaobaoUniversalbpStdcategoryFindlistAPIResponse() *TaobaoUniversalbpStdcategoryFindlistAPIResponse {
	return poolTaobaoUniversalbpStdcategoryFindlistAPIResponse.Get().(*TaobaoUniversalbpStdcategoryFindlistAPIResponse)
}

// ReleaseTaobaoUniversalbpStdcategoryFindlistAPIResponse 将 TaobaoUniversalbpStdcategoryFindlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpStdcategoryFindlistAPIResponse(v *TaobaoUniversalbpStdcategoryFindlistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpStdcategoryFindlistAPIResponse.Put(v)
}
