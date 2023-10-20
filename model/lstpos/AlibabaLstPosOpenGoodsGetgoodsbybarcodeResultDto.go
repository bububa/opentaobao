package lstpos

import (
	"sync"
)

// AlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto 结构体
type AlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto struct {
	// 错误信息描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 接口具体返回的业务数据对象
	Module *GoodsDto `json:"module,omitempty" xml:"module,omitempty"`
	// 业务错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 接口调用是否成功 true:调用成功 false:调用失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto)
	},
}

// GetAlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto() 从对象池中获取AlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto
func GetAlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto() *AlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto {
	return poolAlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto.Get().(*AlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto)
}

// ReleaseAlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto 释放AlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto
func ReleaseAlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto(v *AlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto) {
	v.ErrorMessage = ""
	v.Module = nil
	v.ErrorCode = 0
	v.Success = false
	poolAlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto.Put(v)
}
