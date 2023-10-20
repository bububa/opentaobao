package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemUpdateAPIResponse 根据商品ID或商家编码修改后端商品 API返回值
// taobao.scitem.update
//
// 根据商品ID或商家编码修改后端商品
type TaobaoScitemUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoScitemUpdateAPIResponseModel
}

// TaobaoScitemUpdateAPIResponseModel is 根据商品ID或商家编码修改后端商品 成功返回结果
type TaobaoScitemUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新商品数量,1表示成功更新了一条数据，0：表示未找到匹配的数据
	UpdateRows int64 `json:"update_rows,omitempty" xml:"update_rows,omitempty"`
}
