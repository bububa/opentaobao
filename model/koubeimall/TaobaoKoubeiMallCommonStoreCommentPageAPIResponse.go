package koubeimall

import (
	"encoding/xml"

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

// TaobaoKoubeiMallCommonStoreCommentPageAPIResponseModel is 分页查询门店评论详情信息 成功返回结果
type TaobaoKoubeiMallCommonStoreCommentPageAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_store_comment_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiMallCommonStoreCommentPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
