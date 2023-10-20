package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductskuupdateAPIResponse 产品sku编辑接口 API返回值
// taobao.fenxiao.product.sku.update
//
// 产品SKU信息更新
type TaobaofenxiaoproductskuupdateAPIResponse struct {
	model.CommonResponse
	TaobaofenxiaoproductskuupdateAPIResponseModel
}

// TaobaofenxiaoproductskuupdateAPIResponseModel is 产品sku编辑接口 成功返回结果
type TaobaofenxiaoproductskuupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_sku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
