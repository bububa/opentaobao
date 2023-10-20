package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse 分配joinCode API返回值
// alibaba.cloudgame.interactive.game.joincode.assign
//
// 分配joinCode
type AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse struct {
	model.CommonResponse
	AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponseModel).Reset()
}

// AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponseModel is 分配joinCode 成功返回结果
type AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_joincode_assign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaCloudgameInteractiveGameJoincodeAssignResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse)
	},
}

// GetAlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse
func GetAlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse() *AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse {
	return poolAlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse.Get().(*AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse)
}

// ReleaseAlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse 将 AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse(v *AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse.Put(v)
}
