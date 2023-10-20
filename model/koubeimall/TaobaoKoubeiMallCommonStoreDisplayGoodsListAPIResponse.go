package koubeimall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonstoredisplaygoodslistAPIResponse 查询门店推荐菜信息 API返回值
// taobao.koubei.mall.common.store.display.goods.list
//
// 提供查询口碑商圈内的门店推荐菜信息
type TaobaokoubeimallcommonstoredisplaygoodslistAPIResponse struct {
	model.CommonResponse
	TaobaokoubeimallcommonstoredisplaygoodslistAPIResponseModel
}

// TaobaokoubeimallcommonstoredisplaygoodslistAPIResponseModel is 查询门店推荐菜信息 成功返回结果
type TaobaokoubeimallcommonstoredisplaygoodslistAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_store_display_goods_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// API接口返回的result模型
	Result *TaobaokoubeimallcommonstoredisplaygoodslistResult `json:"result,omitempty" xml:"result,omitempty"`
}
