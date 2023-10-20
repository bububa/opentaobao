package pur

import (
	"sync"
)

// ActionResult 结构体
type ActionResult struct {
	// 系统自动生成
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 系统自动生成
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// https://
	RedirectUrl string `json:"redirect_url,omitempty" xml:"redirect_url,omitempty"`
	// 返回提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的产品链接
	ProductUrl string `json:"product_url,omitempty" xml:"product_url,omitempty"`
	// 套餐标识
	PackageId string `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 返回结果
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误码
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 返回值
	RetValue string `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
	// 系统自动生成
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolActionResult = sync.Pool{
	New: func() any {
		return new(ActionResult)
	},
}

// GetActionResult() 从对象池中获取ActionResult
func GetActionResult() *ActionResult {
	return poolActionResult.Get().(*ActionResult)
}

// ReleaseActionResult 释放ActionResult
func ReleaseActionResult(v *ActionResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.RedirectUrl = ""
	v.Message = ""
	v.ProductUrl = ""
	v.PackageId = ""
	v.Content = ""
	v.Error = ""
	v.RetValue = ""
	v.Success = false
	v.IsSuccess = false
	poolActionResult.Put(v)
}
