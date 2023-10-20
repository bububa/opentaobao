package elife

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoElifeLifecardReconAPIResponse 查询对账文件地址接口 API返回值
// taobao.elife.lifecard.recon
//
// 查询对账文件地址接口
type TaobaoElifeLifecardReconAPIResponse struct {
	model.CommonResponse
	TaobaoElifeLifecardReconAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoElifeLifecardReconAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoElifeLifecardReconAPIResponseModel).Reset()
}

// TaobaoElifeLifecardReconAPIResponseModel is 查询对账文件地址接口 成功返回结果
type TaobaoElifeLifecardReconAPIResponseModel struct {
	XMLName xml.Name `xml:"elife_lifecard_recon_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 对账日期
	OpDate string `json:"op_date,omitempty" xml:"op_date,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 下载下载文件
	ReconFileUrl string `json:"recon_file_url,omitempty" xml:"recon_file_url,omitempty"`
	// 成功标志
	Successed bool `json:"successed,omitempty" xml:"successed,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoElifeLifecardReconAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.OpDate = ""
	m.ResultCode = ""
	m.ReconFileUrl = ""
	m.Successed = false
}

var poolTaobaoElifeLifecardReconAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoElifeLifecardReconAPIResponse)
	},
}

// GetTaobaoElifeLifecardReconAPIResponse 从 sync.Pool 获取 TaobaoElifeLifecardReconAPIResponse
func GetTaobaoElifeLifecardReconAPIResponse() *TaobaoElifeLifecardReconAPIResponse {
	return poolTaobaoElifeLifecardReconAPIResponse.Get().(*TaobaoElifeLifecardReconAPIResponse)
}

// ReleaseTaobaoElifeLifecardReconAPIResponse 将 TaobaoElifeLifecardReconAPIResponse 保存到 sync.Pool
func ReleaseTaobaoElifeLifecardReconAPIResponse(v *TaobaoElifeLifecardReconAPIResponse) {
	v.Reset()
	poolTaobaoElifeLifecardReconAPIResponse.Put(v)
}
