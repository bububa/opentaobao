package tbk

import (
	"sync"
)

// TaobaoTbkScPunishOrderGetRpcResult 结构体
type TaobaoTbkScPunishOrderGetRpcResult struct {
	// 业务出错的描述
	BizErrorDesc string `json:"biz_error_desc,omitempty" xml:"biz_error_desc,omitempty"`
	// 执行结果
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果
	Data *PageResult `json:"data,omitempty" xml:"data,omitempty"`
	// 业务出错的状态码
	BizErrorCode int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 执行结果状态码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolTaobaoTbkScPunishOrderGetRpcResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScPunishOrderGetRpcResult)
	},
}

// GetTaobaoTbkScPunishOrderGetRpcResult() 从对象池中获取TaobaoTbkScPunishOrderGetRpcResult
func GetTaobaoTbkScPunishOrderGetRpcResult() *TaobaoTbkScPunishOrderGetRpcResult {
	return poolTaobaoTbkScPunishOrderGetRpcResult.Get().(*TaobaoTbkScPunishOrderGetRpcResult)
}

// ReleaseTaobaoTbkScPunishOrderGetRpcResult 释放TaobaoTbkScPunishOrderGetRpcResult
func ReleaseTaobaoTbkScPunishOrderGetRpcResult(v *TaobaoTbkScPunishOrderGetRpcResult) {
	v.BizErrorDesc = ""
	v.ResultMsg = ""
	v.Data = nil
	v.BizErrorCode = 0
	v.ResultCode = 0
	poolTaobaoTbkScPunishOrderGetRpcResult.Put(v)
}
