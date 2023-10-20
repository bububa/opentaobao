package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoscitemmapaddAPIResponse 创建IC商品与后端商品的映射关系 API返回值
// taobao.scitem.map.add
//
// 创建IC商品或分销商品与后端商品的映射关系
type TaobaoscitemmapaddAPIResponse struct {
	model.CommonResponse
	TaobaoscitemmapaddAPIResponseModel
}

// TaobaoscitemmapaddAPIResponseModel is 创建IC商品与后端商品的映射关系 成功返回结果
type TaobaoscitemmapaddAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_map_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用返回结果信息：商家编码
	OuterCode string `json:"outer_code,omitempty" xml:"outer_code,omitempty"`
}
