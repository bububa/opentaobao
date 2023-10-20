package qimen

import (
	"sync"
)

// Process 结构体
type Process struct {
	// 单据状态(NEW=新增;ACCEPT=仓库接单;PRINT=打印;PICK=捡货;CHECK=复核;PACKAGE=打包;WEIGH=称重;READY=待提货;DELIVERED=已发货;EXCEPTION=异常;CLOSED=关闭;CANCELED=取消;REJECT=仓库拒单;REFUSE=客户拒签;CANCELEDFAIL=取消失败;SIGN=签收;TMSCANCELED=快递拦截;PARTFULFILLED=部分收货完成;FULFILLED=收货完成;PARTDELIVERED=部分发货完成;OTHER=其他;只传英文编码)
	ProcessStatus string `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
	// 该状态操作员编码
	OperatorCode string `json:"operatorCode,omitempty" xml:"operatorCode,omitempty"`
	// 该状态操作员姓名
	OperatorName string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// 该状态操作时间(YYYY-MM-DD HH:MM:SS)
	OperateTime string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// 操作内容
	OperateInfo string `json:"operateInfo,omitempty" xml:"operateInfo,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 运单号
	ExpressCode string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
}

var poolProcess = sync.Pool{
	New: func() any {
		return new(Process)
	},
}

// GetProcess() 从对象池中获取Process
func GetProcess() *Process {
	return poolProcess.Get().(*Process)
}

// ReleaseProcess 释放Process
func ReleaseProcess(v *Process) {
	v.ProcessStatus = ""
	v.OperatorCode = ""
	v.OperatorName = ""
	v.OperateTime = ""
	v.OperateInfo = ""
	v.Remark = ""
	v.ExpressCode = ""
	poolProcess.Put(v)
}
