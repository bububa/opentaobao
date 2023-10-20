package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscoptimusmaterialAPIResponse 淘宝客-服务商-物料精选 API返回值
// taobao.tbk.sc.optimus.material
//
// 服务商使用。支持入参推广者对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广者对应的推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaotbkscoptimusmaterialAPIResponse struct {
	model.CommonResponse
	TaobaotbkscoptimusmaterialAPIResponseModel
}

// TaobaotbkscoptimusmaterialAPIResponseModel is 淘宝客-服务商-物料精选 成功返回结果
type TaobaotbkscoptimusmaterialAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_optimus_material_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	ResultList []TaobaotbkscoptimusmaterialMapData `json:"result_list,omitempty" xml:"result_list>taobaotbkscoptimusmaterial_map_data,omitempty"`
	// 推荐信息-是否抄底
	IsDefault string `json:"is_default,omitempty" xml:"is_default,omitempty"`
}
