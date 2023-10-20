package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpShopcategoryFindlistAPIResponse 人群相关类目查询 API返回值
// taobao.universalbp.shopcategory.findlist
//
// 查询店铺所属的类目信息
type TaobaoUniversalbpShopcategoryFindlistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpShopcategoryFindlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpShopcategoryFindlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpShopcategoryFindlistAPIResponseModel).Reset()
}

// TaobaoUniversalbpShopcategoryFindlistAPIResponseModel is 人群相关类目查询 成功返回结果
type TaobaoUniversalbpShopcategoryFindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_shopcategory_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpShopcategoryFindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpShopcategoryFindlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpShopcategoryFindlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpShopcategoryFindlistAPIResponse)
	},
}

// GetTaobaoUniversalbpShopcategoryFindlistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpShopcategoryFindlistAPIResponse
func GetTaobaoUniversalbpShopcategoryFindlistAPIResponse() *TaobaoUniversalbpShopcategoryFindlistAPIResponse {
	return poolTaobaoUniversalbpShopcategoryFindlistAPIResponse.Get().(*TaobaoUniversalbpShopcategoryFindlistAPIResponse)
}

// ReleaseTaobaoUniversalbpShopcategoryFindlistAPIResponse 将 TaobaoUniversalbpShopcategoryFindlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpShopcategoryFindlistAPIResponse(v *TaobaoUniversalbpShopcategoryFindlistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpShopcategoryFindlistAPIResponse.Put(v)
}
