package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServicedefinitionQueryskuResult 结构体
type AlibabaSscSupplyplatformServicedefinitionQueryskuResult struct {
	// 服务sku列表
	ServiceSkus []ServiceSkuDto `json:"service_skus,omitempty" xml:"service_skus>service_sku_dto,omitempty"`
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscSupplyplatformServicedefinitionQueryskuResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServicedefinitionQueryskuResult)
	},
}

// GetAlibabaSscSupplyplatformServicedefinitionQueryskuResult() 从对象池中获取AlibabaSscSupplyplatformServicedefinitionQueryskuResult
func GetAlibabaSscSupplyplatformServicedefinitionQueryskuResult() *AlibabaSscSupplyplatformServicedefinitionQueryskuResult {
	return poolAlibabaSscSupplyplatformServicedefinitionQueryskuResult.Get().(*AlibabaSscSupplyplatformServicedefinitionQueryskuResult)
}

// ReleaseAlibabaSscSupplyplatformServicedefinitionQueryskuResult 释放AlibabaSscSupplyplatformServicedefinitionQueryskuResult
func ReleaseAlibabaSscSupplyplatformServicedefinitionQueryskuResult(v *AlibabaSscSupplyplatformServicedefinitionQueryskuResult) {
	v.ServiceSkus = v.ServiceSkus[:0]
	v.DisplayMsg = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = ""
	poolAlibabaSscSupplyplatformServicedefinitionQueryskuResult.Put(v)
}
