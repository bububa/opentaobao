package koubeimall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonMallAuthPageAPIResponse 分页查询已授权的商圈列表信息 API返回值
// taobao.koubei.mall.common.mall.auth.page
//
// 分页查询口碑已授权商圈的列表信息
type TaobaoKoubeiMallCommonMallAuthPageAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonMallAuthPageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonMallAuthPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiMallCommonMallAuthPageAPIResponseModel).Reset()
}

// TaobaoKoubeiMallCommonMallAuthPageAPIResponseModel is 分页查询已授权的商圈列表信息 成功返回结果
type TaobaoKoubeiMallCommonMallAuthPageAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_mall_auth_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// API接口返回的result模型
	Result *TaobaoKoubeiMallCommonMallAuthPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonMallAuthPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiMallCommonMallAuthPageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonMallAuthPageAPIResponse)
	},
}

// GetTaobaoKoubeiMallCommonMallAuthPageAPIResponse 从 sync.Pool 获取 TaobaoKoubeiMallCommonMallAuthPageAPIResponse
func GetTaobaoKoubeiMallCommonMallAuthPageAPIResponse() *TaobaoKoubeiMallCommonMallAuthPageAPIResponse {
	return poolTaobaoKoubeiMallCommonMallAuthPageAPIResponse.Get().(*TaobaoKoubeiMallCommonMallAuthPageAPIResponse)
}

// ReleaseTaobaoKoubeiMallCommonMallAuthPageAPIResponse 将 TaobaoKoubeiMallCommonMallAuthPageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiMallCommonMallAuthPageAPIResponse(v *TaobaoKoubeiMallCommonMallAuthPageAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiMallCommonMallAuthPageAPIResponse.Put(v)
}
