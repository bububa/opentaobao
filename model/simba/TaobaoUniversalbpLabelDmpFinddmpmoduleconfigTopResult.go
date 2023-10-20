package simba

import (
	"sync"
)

// TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult 结构体
type TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	DmpModuleConfigVO *DmpModuleConfigVo `json:"dmp_module_config_v_o,omitempty" xml:"dmp_module_config_v_o,omitempty"`
}

var poolTaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult)
	},
}

// GetTaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult() 从对象池中获取TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult
func GetTaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult() *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult {
	return poolTaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult.Get().(*TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult)
}

// ReleaseTaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult 释放TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult
func ReleaseTaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult(v *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult) {
	v.Info = nil
	v.DmpModuleConfigVO = nil
	poolTaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult.Put(v)
}
