package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeOrderstatusSyncAPIResponse 线下门店派单以及单据相关操作接口 API返回值
// taobao.jst.astrolabe.orderstatus.sync
//
// 针对ERP系统部署在门店的商家，将派单状态回流到星盘
type TaobaoJstAstrolabeOrderstatusSyncAPIResponse struct {
	model.CommonResponse
	TaobaoJstAstrolabeOrderstatusSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeOrderstatusSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstAstrolabeOrderstatusSyncAPIResponseModel).Reset()
}

// TaobaoJstAstrolabeOrderstatusSyncAPIResponseModel is 线下门店派单以及单据相关操作接口 成功返回结果
type TaobaoJstAstrolabeOrderstatusSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_orderstatus_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeOrderstatusSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.Message = ""
}

var poolTaobaoJstAstrolabeOrderstatusSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeOrderstatusSyncAPIResponse)
	},
}

// GetTaobaoJstAstrolabeOrderstatusSyncAPIResponse 从 sync.Pool 获取 TaobaoJstAstrolabeOrderstatusSyncAPIResponse
func GetTaobaoJstAstrolabeOrderstatusSyncAPIResponse() *TaobaoJstAstrolabeOrderstatusSyncAPIResponse {
	return poolTaobaoJstAstrolabeOrderstatusSyncAPIResponse.Get().(*TaobaoJstAstrolabeOrderstatusSyncAPIResponse)
}

// ReleaseTaobaoJstAstrolabeOrderstatusSyncAPIResponse 将 TaobaoJstAstrolabeOrderstatusSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstAstrolabeOrderstatusSyncAPIResponse(v *TaobaoJstAstrolabeOrderstatusSyncAPIResponse) {
	v.Reset()
	poolTaobaoJstAstrolabeOrderstatusSyncAPIResponse.Put(v)
}
