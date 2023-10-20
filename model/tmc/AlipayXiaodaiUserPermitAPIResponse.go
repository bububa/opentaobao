package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlipayXiaodaiUserPermitAPIResponse 阿里金融用户授权 API返回值
// alipay.xiaodai.user.permit
//
// 阿里金融为用户开通消息通道接口
type AlipayXiaodaiUserPermitAPIResponse struct {
	model.CommonResponse
	AlipayXiaodaiUserPermitAPIResponseModel
}

// Reset 清空结构体
func (m *AlipayXiaodaiUserPermitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlipayXiaodaiUserPermitAPIResponseModel).Reset()
}

// AlipayXiaodaiUserPermitAPIResponseModel is 阿里金融用户授权 成功返回结果
type AlipayXiaodaiUserPermitAPIResponseModel struct {
	XMLName xml.Name `xml:"alipay_xiaodai_user_permit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlipayXiaodaiUserPermitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlipayXiaodaiUserPermitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlipayXiaodaiUserPermitAPIResponse)
	},
}

// GetAlipayXiaodaiUserPermitAPIResponse 从 sync.Pool 获取 AlipayXiaodaiUserPermitAPIResponse
func GetAlipayXiaodaiUserPermitAPIResponse() *AlipayXiaodaiUserPermitAPIResponse {
	return poolAlipayXiaodaiUserPermitAPIResponse.Get().(*AlipayXiaodaiUserPermitAPIResponse)
}

// ReleaseAlipayXiaodaiUserPermitAPIResponse 将 AlipayXiaodaiUserPermitAPIResponse 保存到 sync.Pool
func ReleaseAlipayXiaodaiUserPermitAPIResponse(v *AlipayXiaodaiUserPermitAPIResponse) {
	v.Reset()
	poolAlipayXiaodaiUserPermitAPIResponse.Put(v)
}
