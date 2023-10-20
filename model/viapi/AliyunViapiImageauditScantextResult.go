package viapi

import (
	"sync"
)

// AliyunViapiImageauditScantextResult 结构体
type AliyunViapiImageauditScantextResult struct {
	// 单个文本的检测结果
	Details []AliyunViapiImageauditScantextDetail `json:"details,omitempty" xml:"details>aliyun_viapi_imageaudit_scantext_detail,omitempty"`
	// 建议用户执行的操作，取值范围：  pass：文本正常 review：需要人工审核 block：文本违规，可以直接删除或者做限制处理
	Suggestion string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
	// 检测结果的分类
	Label string `json:"label,omitempty" xml:"label,omitempty"`
	// 结果为该分类的概率，取值范围为[0.00-100.00]。值越高，表示越有可能属于该分类。  说明 分值仅供参考，您需要关注label和suggestion内容
	Rate int64 `json:"rate,omitempty" xml:"rate,omitempty"`
}

var poolAliyunViapiImageauditScantextResult = sync.Pool{
	New: func() any {
		return new(AliyunViapiImageauditScantextResult)
	},
}

// GetAliyunViapiImageauditScantextResult() 从对象池中获取AliyunViapiImageauditScantextResult
func GetAliyunViapiImageauditScantextResult() *AliyunViapiImageauditScantextResult {
	return poolAliyunViapiImageauditScantextResult.Get().(*AliyunViapiImageauditScantextResult)
}

// ReleaseAliyunViapiImageauditScantextResult 释放AliyunViapiImageauditScantextResult
func ReleaseAliyunViapiImageauditScantextResult(v *AliyunViapiImageauditScantextResult) {
	v.Details = v.Details[:0]
	v.Suggestion = ""
	v.Label = ""
	v.Rate = 0
	poolAliyunViapiImageauditScantextResult.Put(v)
}
