package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgoptimuspromotionAPIResponse 淘宝客-推广者-权益物料精选 API返回值
// taobao.tbk.dg.optimus.promotion
//
// 推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
type TaobaotbkdgoptimuspromotionAPIResponse struct {
	model.CommonResponse
	TaobaotbkdgoptimuspromotionAPIResponseModel
}

// TaobaotbkdgoptimuspromotionAPIResponseModel is 淘宝客-推广者-权益物料精选 成功返回结果
type TaobaotbkdgoptimuspromotionAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_optimus_promotion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	ResultList []TaobaotbkdgoptimuspromotionMapData `json:"result_list,omitempty" xml:"result_list>taobaotbkdgoptimuspromotion_map_data,omitempty"`
}
