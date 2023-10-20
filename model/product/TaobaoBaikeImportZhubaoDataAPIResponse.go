package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaikeImportZhubaoDataAPIResponse 导入数据到商品百科服务 API返回值
// taobao.baike.import.zhubao.data
//
// 用于接入外部数据录入到商品百科中
type TaobaoBaikeImportZhubaoDataAPIResponse struct {
	model.CommonResponse
	TaobaoBaikeImportZhubaoDataAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaikeImportZhubaoDataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaikeImportZhubaoDataAPIResponseModel).Reset()
}

// TaobaoBaikeImportZhubaoDataAPIResponseModel is 导入数据到商品百科服务 成功返回结果
type TaobaoBaikeImportZhubaoDataAPIResponseModel struct {
	XMLName xml.Name `xml:"baike_import_zhubao_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoBaikeImportZhubaoDataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaikeImportZhubaoDataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBaikeImportZhubaoDataAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaikeImportZhubaoDataAPIResponse)
	},
}

// GetTaobaoBaikeImportZhubaoDataAPIResponse 从 sync.Pool 获取 TaobaoBaikeImportZhubaoDataAPIResponse
func GetTaobaoBaikeImportZhubaoDataAPIResponse() *TaobaoBaikeImportZhubaoDataAPIResponse {
	return poolTaobaoBaikeImportZhubaoDataAPIResponse.Get().(*TaobaoBaikeImportZhubaoDataAPIResponse)
}

// ReleaseTaobaoBaikeImportZhubaoDataAPIResponse 将 TaobaoBaikeImportZhubaoDataAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaikeImportZhubaoDataAPIResponse(v *TaobaoBaikeImportZhubaoDataAPIResponse) {
	v.Reset()
	poolTaobaoBaikeImportZhubaoDataAPIResponse.Put(v)
}
