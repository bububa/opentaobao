package koubeimall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonStoreCommentPageAPIResponse 分页查询门店评论详情信息 API返回值
// taobao.koubei.mall.common.store.comment.page
//
// 查询口碑综合体内的门店评论信息
type TaobaoKoubeiMallCommonStoreCommentPageAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonStoreCommentPageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonStoreCommentPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiMallCommonStoreCommentPageAPIResponseModel).Reset()
}

// TaobaoKoubeiMallCommonStoreCommentPageAPIResponseModel is 分页查询门店评论详情信息 成功返回结果
type TaobaoKoubeiMallCommonStoreCommentPageAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_store_comment_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiMallCommonStoreCommentPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonStoreCommentPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiMallCommonStoreCommentPageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonStoreCommentPageAPIResponse)
	},
}

// GetTaobaoKoubeiMallCommonStoreCommentPageAPIResponse 从 sync.Pool 获取 TaobaoKoubeiMallCommonStoreCommentPageAPIResponse
func GetTaobaoKoubeiMallCommonStoreCommentPageAPIResponse() *TaobaoKoubeiMallCommonStoreCommentPageAPIResponse {
	return poolTaobaoKoubeiMallCommonStoreCommentPageAPIResponse.Get().(*TaobaoKoubeiMallCommonStoreCommentPageAPIResponse)
}

// ReleaseTaobaoKoubeiMallCommonStoreCommentPageAPIResponse 将 TaobaoKoubeiMallCommonStoreCommentPageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiMallCommonStoreCommentPageAPIResponse(v *TaobaoKoubeiMallCommonStoreCommentPageAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiMallCommonStoreCommentPageAPIResponse.Put(v)
}
