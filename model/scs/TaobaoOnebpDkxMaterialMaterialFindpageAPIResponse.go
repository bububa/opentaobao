package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxMaterialMaterialFindpageAPIResponse 获取商品池 API返回值
// taobao.onebp.dkx.material.material.findpage
//
// 获取商品池。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa。
type TaobaoOnebpDkxMaterialMaterialFindpageAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxMaterialMaterialFindpageAPIResponseModel
}

// TaobaoOnebpDkxMaterialMaterialFindpageAPIResponseModel is 获取商品池 成功返回结果
type TaobaoOnebpDkxMaterialMaterialFindpageAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_material_material_findpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxMaterialMaterialFindpageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
