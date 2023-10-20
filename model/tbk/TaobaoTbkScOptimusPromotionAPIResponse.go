package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscoptimuspromotionAPIResponse 淘宝客-服务商-权益物料精选 API返回值
// taobao.tbk.sc.optimus.promotion
//
// 服务商使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
type TaobaotbkscoptimuspromotionAPIResponse struct {
	model.CommonResponse
	TaobaotbkscoptimuspromotionAPIResponseModel
}

// TaobaotbkscoptimuspromotionAPIResponseModel is 淘宝客-服务商-权益物料精选 成功返回结果
type TaobaotbkscoptimuspromotionAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_optimus_promotion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	Resultlist []TaobaotbkscoptimuspromotionMapData `json:"result_list,omitempty" xml:"result_list>taobaotbkscoptimuspromotion_map_data,omitempty"`
}
