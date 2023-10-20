package cloudgame

import (
	"sync"
)

// AlibabaCloudgameInteractiveGameJoincodeAssignResult 结构体
type AlibabaCloudgameInteractiveGameJoincodeAssignResult struct {
	// 返回状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *JoinCodeAssignResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameInteractiveGameJoincodeAssignResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameJoincodeAssignResult)
	},
}

// GetAlibabaCloudgameInteractiveGameJoincodeAssignResult() 从对象池中获取AlibabaCloudgameInteractiveGameJoincodeAssignResult
func GetAlibabaCloudgameInteractiveGameJoincodeAssignResult() *AlibabaCloudgameInteractiveGameJoincodeAssignResult {
	return poolAlibabaCloudgameInteractiveGameJoincodeAssignResult.Get().(*AlibabaCloudgameInteractiveGameJoincodeAssignResult)
}

// ReleaseAlibabaCloudgameInteractiveGameJoincodeAssignResult 释放AlibabaCloudgameInteractiveGameJoincodeAssignResult
func ReleaseAlibabaCloudgameInteractiveGameJoincodeAssignResult(v *AlibabaCloudgameInteractiveGameJoincodeAssignResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameInteractiveGameJoincodeAssignResult.Put(v)
}
