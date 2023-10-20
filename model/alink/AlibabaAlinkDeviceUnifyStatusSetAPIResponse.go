package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceUnifyStatusSetAPIResponse 设置设备标准属性状态 API返回值
// alibaba.alink.device.unify.status.set
//
// 操作用户绑定的设备
type AlibabaAlinkDeviceUnifyStatusSetAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkDeviceUnifyStatusSetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceUnifyStatusSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkDeviceUnifyStatusSetAPIResponseModel).Reset()
}

// AlibabaAlinkDeviceUnifyStatusSetAPIResponseModel is 设置设备标准属性状态 成功返回结果
type AlibabaAlinkDeviceUnifyStatusSetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_device_unify_status_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceUnifyStatusSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkDeviceUnifyStatusSetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkDeviceUnifyStatusSetAPIResponse)
	},
}

// GetAlibabaAlinkDeviceUnifyStatusSetAPIResponse 从 sync.Pool 获取 AlibabaAlinkDeviceUnifyStatusSetAPIResponse
func GetAlibabaAlinkDeviceUnifyStatusSetAPIResponse() *AlibabaAlinkDeviceUnifyStatusSetAPIResponse {
	return poolAlibabaAlinkDeviceUnifyStatusSetAPIResponse.Get().(*AlibabaAlinkDeviceUnifyStatusSetAPIResponse)
}

// ReleaseAlibabaAlinkDeviceUnifyStatusSetAPIResponse 将 AlibabaAlinkDeviceUnifyStatusSetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkDeviceUnifyStatusSetAPIResponse(v *AlibabaAlinkDeviceUnifyStatusSetAPIResponse) {
	v.Reset()
	poolAlibabaAlinkDeviceUnifyStatusSetAPIResponse.Put(v)
}
