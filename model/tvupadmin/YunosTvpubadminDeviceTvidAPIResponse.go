package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindevicetvidAPIResponse 查询终端信息 API返回值
// yunos.tvpubadmin.device.tvid
//
// 通过UUID查询终端信息
type YunostvpubadmindevicetvidAPIResponse struct {
	model.CommonResponse
	YunostvpubadmindevicetvidAPIResponseModel
}

// YunostvpubadmindevicetvidAPIResponseModel is 查询终端信息 成功返回结果
type YunostvpubadmindevicetvidAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_tvid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *DeviceInfoDo `json:"object,omitempty" xml:"object,omitempty"`
}
