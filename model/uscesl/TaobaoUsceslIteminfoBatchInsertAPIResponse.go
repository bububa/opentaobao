package uscesl

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouscesliteminfobatchinsertAPIResponse 按商家批量写入商品接口 API返回值
// taobao.uscesl.iteminfo.batch.insert
//
// 【电子价签】支持按照商家-门店维度批量写入商品数据
type TaobaouscesliteminfobatchinsertAPIResponse struct {
	model.CommonResponse
	TaobaouscesliteminfobatchinsertAPIResponseModel
}

// TaobaouscesliteminfobatchinsertAPIResponseModel is 按商家批量写入商品接口 成功返回结果
type TaobaouscesliteminfobatchinsertAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_iteminfo_batch_insert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
