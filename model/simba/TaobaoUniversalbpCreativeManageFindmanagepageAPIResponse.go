package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse 创意库查询创意列表 API返回值
// taobao.universalbp.creative.manage.findmanagepage
//
// 创意库查询创意列表
type TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCreativeManageFindmanagepageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpCreativeManageFindmanagepageAPIResponseModel).Reset()
}

// TaobaoUniversalbpCreativeManageFindmanagepageAPIResponseModel is 创意库查询创意列表 成功返回结果
type TaobaoUniversalbpCreativeManageFindmanagepageAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_creative_manage_findmanagepage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCreativeManageFindmanagepageTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCreativeManageFindmanagepageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpCreativeManageFindmanagepageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse)
	},
}

// GetTaobaoUniversalbpCreativeManageFindmanagepageAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse
func GetTaobaoUniversalbpCreativeManageFindmanagepageAPIResponse() *TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse {
	return poolTaobaoUniversalbpCreativeManageFindmanagepageAPIResponse.Get().(*TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse)
}

// ReleaseTaobaoUniversalbpCreativeManageFindmanagepageAPIResponse 将 TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpCreativeManageFindmanagepageAPIResponse(v *TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpCreativeManageFindmanagepageAPIResponse.Put(v)
}
