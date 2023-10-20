package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// OdpsAliyuncsComDescribeOdpsService20141215APIResponse 查询ODPS服务 API返回值
// odps.aliyuncs.com.DescribeOdpsService.2014-12-15
//
// 查询ODPS服务
type OdpsAliyuncsComDescribeOdpsService20141215APIResponse struct {
	model.CommonResponse
	OdpsAliyuncsComDescribeOdpsService20141215APIResponseModel
}

// Reset 清空结构体
func (m *OdpsAliyuncsComDescribeOdpsService20141215APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.OdpsAliyuncsComDescribeOdpsService20141215APIResponseModel).Reset()
}

// OdpsAliyuncsComDescribeOdpsService20141215APIResponseModel is 查询ODPS服务 成功返回结果
type OdpsAliyuncsComDescribeOdpsService20141215APIResponseModel struct {
	XMLName xml.Name `xml:"odps_aliyuncs_com_DescribeOdpsService_2014-12-15_response"`
	// 请求ID
	RequestId string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务Id
	ServiceId string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// 开通服务时间，ISO 8601时间格式
	OpeningTime string `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	// 业务锁定状态，例如：欠费，安全等。
	OperationLocks string `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty"`
}

// Reset 清空结构体
func (m *OdpsAliyuncsComDescribeOdpsService20141215APIResponseModel) Reset() {
	m.RequestId = ""
	m.RequestId = ""
	m.ServiceId = ""
	m.OpeningTime = ""
	m.OperationLocks = ""
}

var poolOdpsAliyuncsComDescribeOdpsService20141215APIResponse = sync.Pool{
	New: func() any {
		return new(OdpsAliyuncsComDescribeOdpsService20141215APIResponse)
	},
}

// GetOdpsAliyuncsComDescribeOdpsService20141215APIResponse 从 sync.Pool 获取 OdpsAliyuncsComDescribeOdpsService20141215APIResponse
func GetOdpsAliyuncsComDescribeOdpsService20141215APIResponse() *OdpsAliyuncsComDescribeOdpsService20141215APIResponse {
	return poolOdpsAliyuncsComDescribeOdpsService20141215APIResponse.Get().(*OdpsAliyuncsComDescribeOdpsService20141215APIResponse)
}

// ReleaseOdpsAliyuncsComDescribeOdpsService20141215APIResponse 将 OdpsAliyuncsComDescribeOdpsService20141215APIResponse 保存到 sync.Pool
func ReleaseOdpsAliyuncsComDescribeOdpsService20141215APIResponse(v *OdpsAliyuncsComDescribeOdpsService20141215APIResponse) {
	v.Reset()
	poolOdpsAliyuncsComDescribeOdpsService20141215APIResponse.Put(v)
}
