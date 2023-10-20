package dt

import (
	"sync"
)

// RtItemPriceTagBackParam 结构体
type RtItemPriceTagBackParam struct {
	// 数据列表
	PriceTagParamList []RtItemResearchPriceParam `json:"price_tag_param_list,omitempty" xml:"price_tag_param_list>rt_item_research_price_param,omitempty"`
	// 业务编码
	BusiCode string `json:"busi_code,omitempty" xml:"busi_code,omitempty"`
	// 业务来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
}

var poolRtItemPriceTagBackParam = sync.Pool{
	New: func() any {
		return new(RtItemPriceTagBackParam)
	},
}

// GetRtItemPriceTagBackParam() 从对象池中获取RtItemPriceTagBackParam
func GetRtItemPriceTagBackParam() *RtItemPriceTagBackParam {
	return poolRtItemPriceTagBackParam.Get().(*RtItemPriceTagBackParam)
}

// ReleaseRtItemPriceTagBackParam 释放RtItemPriceTagBackParam
func ReleaseRtItemPriceTagBackParam(v *RtItemPriceTagBackParam) {
	v.PriceTagParamList = v.PriceTagParamList[:0]
	v.BusiCode = ""
	v.Source = ""
	poolRtItemPriceTagBackParam.Put(v)
}
