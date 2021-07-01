package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoScitemMapDeleteAPIResponse
失效指定用户的商品与后端商品的映射关系 API返回值
taobao.scitem.map.delete

根据后端商品Id，失效指定用户的商品与后端商品的映射关系 */
type TaobaoScitemMapDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoScitemMapDeleteAPIResponseModel
}

// TaobaoScitemMapDeleteAPIResponseModel is 失效指定用户的商品与后端商品的映射关系 成功返回结果
type TaobaoScitemMapDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_map_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失效条数
	Module int64 `json:"module,omitempty" xml:"module,omitempty"`
}
