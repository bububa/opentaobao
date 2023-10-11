package mtopopen

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
