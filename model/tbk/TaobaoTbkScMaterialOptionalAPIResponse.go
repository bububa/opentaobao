package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScMaterialOptionalAPIResponse 淘宝客-服务商-物料搜索 API返回值
// taobao.tbk.sc.material.optional
//
// 服务商使用。支持入参推广者对应的“推广位”、关键词和相关筛选条件，获取对应的物料信息和推广者对应的推广链接。
type TaobaoTbkScMaterialOptionalAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScMaterialOptionalAPIResponseModel
}

// TaobaoTbkScMaterialOptionalAPIResponseModel is 淘宝客-服务商-物料搜索 成功返回结果
type TaobaoTbkScMaterialOptionalAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_material_optional_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	ResultList []TaobaoTbkScMaterialOptionalMapData `json:"result_list,omitempty" xml:"result_list>taobao_tbk_sc_material_optional_map_data,omitempty"`
	// 本地化-lbs唯一分页标示，请在翻页时作为入参透传
	PageResultKey string `json:"page_result_key,omitempty" xml:"page_result_key,omitempty"`
	// 搜索到符合条件的结果总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
