package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductToChannelImportAPIResponse 产品导入到渠道 API返回值
// taobao.fenxiao.product.to.channel.import
//
// 支持供应商将已有产品导入到某个渠道销售
type TaobaoFenxiaoProductToChannelImportAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductToChannelImportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductToChannelImportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductToChannelImportAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductToChannelImportAPIResponseModel is 产品导入到渠道 成功返回结果
type TaobaoFenxiaoProductToChannelImportAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_to_channel_import_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductToChannelImportAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolTaobaoFenxiaoProductToChannelImportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductToChannelImportAPIResponse)
	},
}

// GetTaobaoFenxiaoProductToChannelImportAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductToChannelImportAPIResponse
func GetTaobaoFenxiaoProductToChannelImportAPIResponse() *TaobaoFenxiaoProductToChannelImportAPIResponse {
	return poolTaobaoFenxiaoProductToChannelImportAPIResponse.Get().(*TaobaoFenxiaoProductToChannelImportAPIResponse)
}

// ReleaseTaobaoFenxiaoProductToChannelImportAPIResponse 将 TaobaoFenxiaoProductToChannelImportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductToChannelImportAPIResponse(v *TaobaoFenxiaoProductToChannelImportAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductToChannelImportAPIResponse.Put(v)
}
