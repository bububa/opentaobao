package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgoptimusmaterialAPIResponse 淘宝客-推广者-物料精选 API返回值
// taobao.tbk.dg.optimus.material
//
// 支持入参对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaotbkdgoptimusmaterialAPIResponse struct {
	model.CommonResponse
	TaobaotbkdgoptimusmaterialAPIResponseModel
}

// TaobaotbkdgoptimusmaterialAPIResponseModel is 淘宝客-推广者-物料精选 成功返回结果
type TaobaotbkdgoptimusmaterialAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_optimus_material_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	Resultlist []TaobaotbkdgoptimusmaterialMapData `json:"result_list,omitempty" xml:"result_list>taobaotbkdgoptimusmaterial_map_data,omitempty"`
	// 推荐信息-是否抄底
	Isdefault string `json:"is_default,omitempty" xml:"is_default,omitempty"`
	// 商品总数-目前只有全品库商品查询有该字段
	Totalcount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
