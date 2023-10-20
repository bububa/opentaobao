package lstwarehouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstBranddatashareStockdataQueryAPIResponse 查询品牌商品实仓库存/周转效能 API返回值
// alibaba.lst.branddatashare.stockdata.query
//
// 品牌商查询授权供应商实仓库存数据
type AlibabaLstBranddatashareStockdataQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstBranddatashareStockdataQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstBranddatashareStockdataQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstBranddatashareStockdataQueryAPIResponseModel).Reset()
}

// AlibabaLstBranddatashareStockdataQueryAPIResponseModel is 查询品牌商品实仓库存/周转效能 成功返回结果
type AlibabaLstBranddatashareStockdataQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_branddatashare_stockdata_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstBranddatashareStockdataQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstBranddatashareStockdataQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstBranddatashareStockdataQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstBranddatashareStockdataQueryAPIResponse)
	},
}

// GetAlibabaLstBranddatashareStockdataQueryAPIResponse 从 sync.Pool 获取 AlibabaLstBranddatashareStockdataQueryAPIResponse
func GetAlibabaLstBranddatashareStockdataQueryAPIResponse() *AlibabaLstBranddatashareStockdataQueryAPIResponse {
	return poolAlibabaLstBranddatashareStockdataQueryAPIResponse.Get().(*AlibabaLstBranddatashareStockdataQueryAPIResponse)
}

// ReleaseAlibabaLstBranddatashareStockdataQueryAPIResponse 将 AlibabaLstBranddatashareStockdataQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstBranddatashareStockdataQueryAPIResponse(v *AlibabaLstBranddatashareStockdataQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstBranddatashareStockdataQueryAPIResponse.Put(v)
}
