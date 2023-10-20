package aedropshiper

import (
	"sync"
)

// AeopNationalQuoteConfiguration 结构体
type AeopNationalQuoteConfiguration struct {
	// 分国家定价规则类型[percentage：基于基准价格按比例配置]
	ConfigurationType string `json:"configuration_type,omitempty" xml:"configuration_type,omitempty"`
	// jsonArray格式的分国家定价规则数据。 1)基于基准价格按比例配置的数据格式：[{&#34;shiptoCountry&#34;:&#34;US&#34;,&#34;percentage&#34;:&#34;5&#34;},{&#34;shiptoCountry&#34;:&#34;RU&#34;,&#34;percentage&#34;:&#34;-2&#34;}] 其中shiptoCountry：ISO两位的国家编码（目前支持11个国家：RU,US,CA,ES,FR,UK,NL,IL,BR,CL,AU）， percentage：相对于基准价的调价比例（百分比整数，支持负数，当前限制&gt;=-30 &amp;&amp; &lt;=100）
	ConfigurationData string `json:"configuration_data,omitempty" xml:"configuration_data,omitempty"`
}

var poolAeopNationalQuoteConfiguration = sync.Pool{
	New: func() any {
		return new(AeopNationalQuoteConfiguration)
	},
}

// GetAeopNationalQuoteConfiguration() 从对象池中获取AeopNationalQuoteConfiguration
func GetAeopNationalQuoteConfiguration() *AeopNationalQuoteConfiguration {
	return poolAeopNationalQuoteConfiguration.Get().(*AeopNationalQuoteConfiguration)
}

// ReleaseAeopNationalQuoteConfiguration 释放AeopNationalQuoteConfiguration
func ReleaseAeopNationalQuoteConfiguration(v *AeopNationalQuoteConfiguration) {
	v.ConfigurationType = ""
	v.ConfigurationData = ""
	poolAeopNationalQuoteConfiguration.Put(v)
}
