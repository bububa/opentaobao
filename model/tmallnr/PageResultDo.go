package tmallnr

import (
	"sync"
)

// PageResultDo 结构体
type PageResultDo struct {
	// 门店集合
	DataList []NrtStoreDto `json:"data_list,omitempty" xml:"data_list>nrt_store_dto,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 数据总数目
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 每页记录数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	TotalPageNum int64 `json:"total_page_num,omitempty" xml:"total_page_num,omitempty"`
	// 当前页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPageResultDo = sync.Pool{
	New: func() any {
		return new(PageResultDo)
	},
}

// GetPageResultDo() 从对象池中获取PageResultDo
func GetPageResultDo() *PageResultDo {
	return poolPageResultDo.Get().(*PageResultDo)
}

// ReleasePageResultDo 释放PageResultDo
func ReleasePageResultDo(v *PageResultDo) {
	v.DataList = v.DataList[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.TotalNum = 0
	v.PageSize = 0
	v.TotalPageNum = 0
	v.PageNum = 0
	v.Success = false
	poolPageResultDo.Put(v)
}
