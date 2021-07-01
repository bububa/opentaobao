package koubeimall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiMallCommonItemDetailQueryAPIResponse
查询商品详情信息 API返回值
taobao.koubei.mall.common.item.detail.query

查询口碑综合体内商品详情信息 */
type TaobaoKoubeiMallCommonItemDetailQueryAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonItemDetailQueryAPIResponseModel
}

// TaobaoKoubeiMallCommonItemDetailQueryAPIResponseModel is 查询商品详情信息 成功返回结果
type TaobaoKoubeiMallCommonItemDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_item_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiMallCommonItemDetailQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
