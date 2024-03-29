package viapi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiGoodstechRecognizeFurniturespuAPIResponse 家居SPU识别 API返回值
// aliyun.viapi.goodstech.recognize.furniturespu
//
// 对输入的家居模型图进行分类，目前类别数可达70类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiGoodstechRecognizeFurniturespuAPIResponse struct {
	model.CommonResponse
	AliyunViapiGoodstechRecognizeFurniturespuAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunViapiGoodstechRecognizeFurniturespuAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiGoodstechRecognizeFurniturespuAPIResponseModel).Reset()
}

// AliyunViapiGoodstechRecognizeFurniturespuAPIResponseModel is 家居SPU识别 成功返回结果
type AliyunViapiGoodstechRecognizeFurniturespuAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_goodstech_recognize_furniturespu_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiGoodstechRecognizeFurniturespuData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliyunViapiGoodstechRecognizeFurniturespuAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiGoodstechRecognizeFurniturespuAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiGoodstechRecognizeFurniturespuAPIResponse)
	},
}

// GetAliyunViapiGoodstechRecognizeFurniturespuAPIResponse 从 sync.Pool 获取 AliyunViapiGoodstechRecognizeFurniturespuAPIResponse
func GetAliyunViapiGoodstechRecognizeFurniturespuAPIResponse() *AliyunViapiGoodstechRecognizeFurniturespuAPIResponse {
	return poolAliyunViapiGoodstechRecognizeFurniturespuAPIResponse.Get().(*AliyunViapiGoodstechRecognizeFurniturespuAPIResponse)
}

// ReleaseAliyunViapiGoodstechRecognizeFurniturespuAPIResponse 将 AliyunViapiGoodstechRecognizeFurniturespuAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiGoodstechRecognizeFurniturespuAPIResponse(v *AliyunViapiGoodstechRecognizeFurniturespuAPIResponse) {
	v.Reset()
	poolAliyunViapiGoodstechRecognizeFurniturespuAPIResponse.Put(v)
}
