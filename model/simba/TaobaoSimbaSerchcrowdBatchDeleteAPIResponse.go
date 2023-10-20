package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSerchcrowdBatchDeleteAPIResponse 单品搜索人群批量取消溢价 API返回值
// taobao.simba.serchcrowd.batch.delete
//
// 删除单品搜索人群溢价功能
type TaobaoSimbaSerchcrowdBatchDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSerchcrowdBatchDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSerchcrowdBatchDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSerchcrowdBatchDeleteAPIResponseModel).Reset()
}

// TaobaoSimbaSerchcrowdBatchDeleteAPIResponseModel is 单品搜索人群批量取消溢价 成功返回结果
type TaobaoSimbaSerchcrowdBatchDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_serchcrowd_batch_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	DeleteList []string `json:"delete_list,omitempty" xml:"delete_list>string,omitempty"`
	// errorDTOList
	ErrorDTOList []string `json:"error_d_t_o_list,omitempty" xml:"error_d_t_o_list>string,omitempty"`
	// key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSerchcrowdBatchDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeleteList = m.DeleteList[:0]
	m.ErrorDTOList = m.ErrorDTOList[:0]
	m.Key = ""
}

var poolTaobaoSimbaSerchcrowdBatchDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSerchcrowdBatchDeleteAPIResponse)
	},
}

// GetTaobaoSimbaSerchcrowdBatchDeleteAPIResponse 从 sync.Pool 获取 TaobaoSimbaSerchcrowdBatchDeleteAPIResponse
func GetTaobaoSimbaSerchcrowdBatchDeleteAPIResponse() *TaobaoSimbaSerchcrowdBatchDeleteAPIResponse {
	return poolTaobaoSimbaSerchcrowdBatchDeleteAPIResponse.Get().(*TaobaoSimbaSerchcrowdBatchDeleteAPIResponse)
}

// ReleaseTaobaoSimbaSerchcrowdBatchDeleteAPIResponse 将 TaobaoSimbaSerchcrowdBatchDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSerchcrowdBatchDeleteAPIResponse(v *TaobaoSimbaSerchcrowdBatchDeleteAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSerchcrowdBatchDeleteAPIResponse.Put(v)
}
