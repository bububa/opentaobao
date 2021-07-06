package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductcatAddAPIResponse 新增产品线 API返回值
// taobao.fenxiao.productcat.add
//
// 新增产品线
type TaobaoFenxiaoProductcatAddAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductcatAddAPIResponseModel
}

// TaobaoFenxiaoProductcatAddAPIResponseModel is 新增产品线 成功返回结果
type TaobaoFenxiaoProductcatAddAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_productcat_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品线ID
	ProductLineId int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
