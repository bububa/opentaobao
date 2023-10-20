package moziacl

import (
	"sync"
)

// AppPermissionPackageResult 结构体
type AppPermissionPackageResult struct {
	// 套权限餐数据列表
	Datas []PermissionPackageEntity `json:"datas,omitempty" xml:"datas>permission_package_entity,omitempty"`
	// 请求唯一id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应message，若失败则返回失败原因
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 扩展字段，与入参扩展字段值相同
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 响应code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 应用下套餐总数
	TotalSize int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	// 每页数据条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 查询第几页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 是否处理成功，成功则为true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAppPermissionPackageResult = sync.Pool{
	New: func() any {
		return new(AppPermissionPackageResult)
	},
}

// GetAppPermissionPackageResult() 从对象池中获取AppPermissionPackageResult
func GetAppPermissionPackageResult() *AppPermissionPackageResult {
	return poolAppPermissionPackageResult.Get().(*AppPermissionPackageResult)
}

// ReleaseAppPermissionPackageResult 释放AppPermissionPackageResult
func ReleaseAppPermissionPackageResult(v *AppPermissionPackageResult) {
	v.Datas = v.Datas[:0]
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.TotalSize = 0
	v.PageSize = 0
	v.CurrentPage = 0
	v.Success = false
	poolAppPermissionPackageResult.Put(v)
}
