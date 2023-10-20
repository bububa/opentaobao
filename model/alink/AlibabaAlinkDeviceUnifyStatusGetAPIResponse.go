package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceUnifyStatusGetAPIResponse 查询设备标准属性最新状态 API返回值
// alibaba.alink.device.unify.status.get
//
// 查询设备最新标准属性状态
type AlibabaAlinkDeviceUnifyStatusGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkDeviceUnifyStatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceUnifyStatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkDeviceUnifyStatusGetAPIResponseModel).Reset()
}

// AlibabaAlinkDeviceUnifyStatusGetAPIResponseModel is 查询设备标准属性最新状态 成功返回结果
type AlibabaAlinkDeviceUnifyStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_device_unify_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceUnifyStatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkDeviceUnifyStatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkDeviceUnifyStatusGetAPIResponse)
	},
}

// GetAlibabaAlinkDeviceUnifyStatusGetAPIResponse 从 sync.Pool 获取 AlibabaAlinkDeviceUnifyStatusGetAPIResponse
func GetAlibabaAlinkDeviceUnifyStatusGetAPIResponse() *AlibabaAlinkDeviceUnifyStatusGetAPIResponse {
	return poolAlibabaAlinkDeviceUnifyStatusGetAPIResponse.Get().(*AlibabaAlinkDeviceUnifyStatusGetAPIResponse)
}

// ReleaseAlibabaAlinkDeviceUnifyStatusGetAPIResponse 将 AlibabaAlinkDeviceUnifyStatusGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkDeviceUnifyStatusGetAPIResponse(v *AlibabaAlinkDeviceUnifyStatusGetAPIResponse) {
	v.Reset()
	poolAlibabaAlinkDeviceUnifyStatusGetAPIResponse.Put(v)
}
