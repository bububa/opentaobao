package mtopopen

import (
	"sync"
)

// QueryPackageListResponse 结构体
type QueryPackageListResponse struct {
	// 包裹信息
	PackageInfos []PackageInfoVo `json:"package_infos,omitempty" xml:"package_infos>package_info_vo,omitempty"`
	// 查询错误码(成功情况无需关注)
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 查询错误原因(成功情况无需关注)
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 包裹总数量
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 查询是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolQueryPackageListResponse = sync.Pool{
	New: func() any {
		return new(QueryPackageListResponse)
	},
}

// GetQueryPackageListResponse() 从对象池中获取QueryPackageListResponse
func GetQueryPackageListResponse() *QueryPackageListResponse {
	return poolQueryPackageListResponse.Get().(*QueryPackageListResponse)
}

// ReleaseQueryPackageListResponse 释放QueryPackageListResponse
func ReleaseQueryPackageListResponse(v *QueryPackageListResponse) {
	v.PackageInfos = v.PackageInfos[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalNum = 0
	v.Success = false
	poolQueryPackageListResponse.Put(v)
}
