package lstwarehouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstBranddatashareSuppliersListQueryAPIResponse 品牌数据授权的供应商列表 API返回值
// alibaba.lst.branddatashare.suppliers.list.query
//
// 品牌商查询品牌数据授权的供应商列表
type AlibabaLstBranddatashareSuppliersListQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstBranddatashareSuppliersListQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstBranddatashareSuppliersListQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstBranddatashareSuppliersListQueryAPIResponseModel).Reset()
}

// AlibabaLstBranddatashareSuppliersListQueryAPIResponseModel is 品牌数据授权的供应商列表 成功返回结果
type AlibabaLstBranddatashareSuppliersListQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_branddatashare_suppliers_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值列表
	ContentList []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstBranddatashareSuppliersListQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ContentList = m.ContentList[:0]
}

var poolAlibabaLstBranddatashareSuppliersListQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstBranddatashareSuppliersListQueryAPIResponse)
	},
}

// GetAlibabaLstBranddatashareSuppliersListQueryAPIResponse 从 sync.Pool 获取 AlibabaLstBranddatashareSuppliersListQueryAPIResponse
func GetAlibabaLstBranddatashareSuppliersListQueryAPIResponse() *AlibabaLstBranddatashareSuppliersListQueryAPIResponse {
	return poolAlibabaLstBranddatashareSuppliersListQueryAPIResponse.Get().(*AlibabaLstBranddatashareSuppliersListQueryAPIResponse)
}

// ReleaseAlibabaLstBranddatashareSuppliersListQueryAPIResponse 将 AlibabaLstBranddatashareSuppliersListQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstBranddatashareSuppliersListQueryAPIResponse(v *AlibabaLstBranddatashareSuppliersListQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstBranddatashareSuppliersListQueryAPIResponse.Put(v)
}
