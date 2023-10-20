package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgOptimusMaterialAPIResponse 淘宝客-推广者-物料精选 API返回值
// taobao.tbk.dg.optimus.material
//
// 支持入参对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaoTbkDgOptimusMaterialAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgOptimusMaterialAPIResponseModel
}

// TaobaoTbkDgOptimusMaterialAPIResponseModel is 淘宝客-推广者-物料精选 成功返回结果
type TaobaoTbkDgOptimusMaterialAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_optimus_material_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	ResultList []TaobaoTbkDgOptimusMaterialMapData `json:"result_list,omitempty" xml:"result_list>taobao_tbk_dg_optimus_material_map_data,omitempty"`
	// 推荐信息-是否抄底
	IsDefault string `json:"is_default,omitempty" xml:"is_default,omitempty"`
	// 商品总数-目前只有全品库商品查询有该字段
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
