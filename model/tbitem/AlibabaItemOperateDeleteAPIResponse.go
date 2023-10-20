package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemOperateDeleteAPIResponse 商品删除 API返回值
// alibaba.item.operate.delete
//
// 商品删除
type AlibabaItemOperateDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaItemOperateDeleteAPIResponseModel
}

// AlibabaItemOperateDeleteAPIResponseModel is 商品删除 成功返回结果
type AlibabaItemOperateDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_operate_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品删除是否成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
