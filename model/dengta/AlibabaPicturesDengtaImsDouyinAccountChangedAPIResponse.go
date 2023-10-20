package dengta

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse 接收发生变化的抖音帐号 API返回值
// alibaba.pictures.dengta.ims.douyin.account.changed
//
// 接收发生变化的抖音帐号
type AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse struct {
	model.CommonResponse
	AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponseModel).Reset()
}

// AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponseModel is 接收发生变化的抖音帐号 成功返回结果
type AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pictures_dengta_ims_douyin_account_changed_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *GeneralResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse)
	},
}

// GetAlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse 从 sync.Pool 获取 AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse
func GetAlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse() *AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse {
	return poolAlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse.Get().(*AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse)
}

// ReleaseAlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse 将 AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse(v *AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse) {
	v.Reset()
	poolAlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse.Put(v)
}
