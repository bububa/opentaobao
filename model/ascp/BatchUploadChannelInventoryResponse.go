package ascp

import (
	"sync"
)

// BatchUploadChannelInventoryResponse 结构体
type BatchUploadChannelInventoryResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务处理结果
	Data *BatchUploadChannelInventoryResult `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBatchUploadChannelInventoryResponse = sync.Pool{
	New: func() any {
		return new(BatchUploadChannelInventoryResponse)
	},
}

// GetBatchUploadChannelInventoryResponse() 从对象池中获取BatchUploadChannelInventoryResponse
func GetBatchUploadChannelInventoryResponse() *BatchUploadChannelInventoryResponse {
	return poolBatchUploadChannelInventoryResponse.Get().(*BatchUploadChannelInventoryResponse)
}

// ReleaseBatchUploadChannelInventoryResponse 释放BatchUploadChannelInventoryResponse
func ReleaseBatchUploadChannelInventoryResponse(v *BatchUploadChannelInventoryResponse) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolBatchUploadChannelInventoryResponse.Put(v)
}
