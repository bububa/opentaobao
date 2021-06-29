package koubeimall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商圈详细信息 API返回值 
taobao.koubei.mall.common.mall.detail.get

查询口碑综合体-商圈详细信息，包含商圈基础信息、门店类目分类、商圈推荐商品等模块信息
*/
type TaobaoKoubeiMallCommonMallDetailGetAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiMallCommonMallDetailGetResponse
}

// 查询商圈详细信息 成功返回结果
type TaobaoKoubeiMallCommonMallDetailGetResponse struct {
    XMLName xml.Name `xml:"koubei_mall_common_mall_detail_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // API接口返回的result模型
    Result   *TaobaoKoubeiMallCommonMallDetailGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
